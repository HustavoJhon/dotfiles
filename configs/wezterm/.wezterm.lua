-- PACMAN.DOTS - WezTerm Configuration

local wezterm = require 'wezterm'
local config = wezterm.config_builder()

config.font = wezterm.font('Hack Nerd Font', { weight = 'Regular' })
config.font_size = 14.0
config.line_height = 1.2

config.color_scheme = 'Catppuccin Mocha'

config.window_background_opacity = 0.85
config.macos_window_background_blur = 30
config.window_decorations = 'RESIZE'
config.window_close_confirmation = 'NeverPrompt'

config.enable_tab_bar = true
config.use_fancy_tab_bar = true
config.tab_bar_at_bottom = false
config.tab_max_width = 32
config.switch_to_last_active_tab_when_stopping_tab = true

config.hide_tab_bar_if_only_one_tab = true

config.audible_bell = 'Disabled'

config.default_prog = { 'zsh' }
config.default_cwd = os.getenv('HOME')

config.keys = {
  { key = 'LeftArrow',  mods = 'CTRL', action = wezterm.action{ ActivatePaneDirection = 'Left' } },
  { key = 'RightArrow', mods = 'CTRL', action = wezterm.action{ ActivatePaneDirection = 'Right' } },
  { key = 'DownArrow',  mods = 'CTRL', action = wezterm.action{ ActivatePaneDirection = 'Down' } },
  { key = 'UpArrow',    mods = 'CTRL', action = wezterm.action{ ActivatePaneDirection = 'Up' } },
  { key = 't', mods = 'CTRL|SHIFT', action = wezterm.action{ SpawnTab = 'CurrentPaneDomain' } },
  { key = 'w', mods = 'CTRL|SHIFT', action = wezterm.action{ CloseCurrentPane = { confirm = true } } },
  { key = 'n', mods = 'CTRL|SHIFT', action = wezterm.action{ SplitVertical = { domain = 'CurrentPaneDomain' } } },
  { key = 'b', mods = 'CTRL|SHIFT', action = wezterm.action{ SplitHorizontal = { domain = 'CurrentPaneDomain' } } },
}

config.colors = {
  tab_bar = {
    active_tab = { bg_color = '#7C3AED', fg_color = '#FFFFFF' },
    inactive_tab = { bg_color = '#1a1b26', fg_color = '#6B7280' },
    inactive_tab_hover = { bg_color = '#2a2b3e', fg_color = '#FFFFFF' },
    new_tab = { bg_color = '#0F0F0F', fg_color = '#6B7280' },
    new_tab_hover = { bg_color = '#7C3AED', fg_color = '#FFFFFF' },
  },
}

return config
