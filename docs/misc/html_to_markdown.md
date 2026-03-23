  Step 1 — JohannesKaufmann/html-to-markdown
  - https://github.com/JohannesKaufmann/html-to-markdown
  - Plugin-based, handles tables/nested lists/rowspan. Actively maintained (v2).

  Step 2 — charmbracelet/glamour
  - https://github.com/charmbracelet/glamour
  - Renders Markdown as styled ANSI terminal output. Powers GitHub CLI and glow. Latest commit March 2026.

  import (
      htmltomd "github.com/JohannesKaufmann/html-to-markdown/v2"
      "github.com/charmbracelet/glamour"
  )

  func renderRSSContent(htmlContent string) (string, error) {
      md, err := htmltomd.ConvertString(htmlContent)
      if err != nil {
          return "", err
      }

      r, err := glamour.NewTermRenderer(
          glamour.WithAutoStyle(),  // detects dark/light terminal
          glamour.WithWordWrap(80),
      )
      if err != nil {
          return "", err
      }

      return r.Render(md)
  }

