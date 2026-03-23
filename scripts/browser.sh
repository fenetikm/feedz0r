#!/usr/bin/env zsh

./fz list --type=posts |
  fzf \
  --preview 'echo {} | sed "s/.*(\([0-9]*\)).*/\1/" | xargs -I {} ./fz post --output=markdown {}'
