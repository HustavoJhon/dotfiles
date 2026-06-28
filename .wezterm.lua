-- PACMAN - WezTerm Configuration
local wezterm = require 'wezterm'

return {
  font = wezterm.font('Hack Nerd Font'),
  font_size = 14.0,
  color_scheme = 'Tokyo Night',
  hide_tab_bar_if_only_one_tab = true,
  window_background_opacity = 0.85,
  macos_window_background_blur = 20,
  window_decorations = 'RESIZE',
  default_cursor_style = 'BlinkingBar',
  keys = {
    { key = 'LeftArrow', mods = 'CTRL', action = wezterm.action.ActivatePaneDirection 'Left' },
    { key = 'RightArrow', mods = 'CTRL', action = wezterm.action.ActivatePaneDirection 'Right' },
    { key = 'UpArrow', mods = 'CTRL', action = wezterm.action.ActivatePaneDirection 'Up' },
    { key = 'DownArrow', mods = 'CTRL', action = wezterm.action.ActivatePaneDirection 'Down' },
  },
}
