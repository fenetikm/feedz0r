Here are concrete improvements based on the actual code:

---

### Bugs / correctness

**1. `watch` kills the loop on any error** (`watch.go:27`)
`log.Fatalln` on a fetch error exits the process. A single bad feed kills the whole watch loop. Should log and continue.

**2. `list` swallows errors** (`list.go:22,30`)
Both error paths do `fmt.Println(...)` then `return nil`. The error is never returned to the caller.

**3. `post` ignores the `glamour.Render` error** (`post.go:49`)
`fmt.Print(out)` runs unconditionally even when `err != nil`, potentially printing garbage. The error is returned after the print.

**4. `add` loses the original DB error** (`add.go:35`)
`errors.New("Couldn't create feed.")` discards the underlying error — you can't tell if it was a duplicate URL or a connection problem.

**5. RSS date parsing is too strict** (`feeds.go:56`)
Only handles `"Mon, 02 Jan 2006 15:04:05 -0700"`. Common variants like `+0000` or `GMT` will silently skip every item in a feed.

---

### Design / maintainability

**6. New `http.Client` on every fetch** (`feeds.go:99`)
`&http.Client{}` is created per-call. A shared client at the state or package level would reuse connections and allow proper configuration (redirect policy, etc.).

**7. `watch` parses a string it just formatted** (`watch.go:15-17`)
`fmt.Sprintf("%dm", mins)` → `time.ParseDuration(...)` is roundabout. Just use `time.Duration(s.Config.Fetch.RefreshMins) * time.Minute`.

**8. No signal handling in `watch`** (`watch.go:24`)
`context.Background()` with no cancellation means Ctrl-C won't cleanly shut down — the ticker is also never stopped (`ticker.Stop()`).

**9. Hardcoded DB path** (`main.go:32`)
`"feedz0r.db"` is hardcoded. Should come from config or an env var so the binary can be run from any working directory.

**10. `help` command is a stub** (`help/help.go`)
`help.Handle` just prints `"Help!"`. Each command already has a `Help()` function defined but they're never wired up. `fz` with no args says `"not enough args"` — it should print usage.

---

### Features worth adding

**11. Atom feed support** — noted as a todo in `feeds.go:20`. Many modern feeds (including the example in `add/add.go`'s help text — `atom.xml`) are Atom, not RSS. Currently they'd parse silently with no items.

**12. New post count on fetch** — `feeds.go:38` has a todo for this. Knowing how many new posts were added vs skipped as duplicates is useful output from `fetch`/`watch`.

**13. `remove` command** — there's no way to delete a feed (or its posts) once added.
