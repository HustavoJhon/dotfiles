-- Cargar packer.nvim
vim.cmd [[packadd packer.nvim]]

-- Configuración de plugins usando packer
require('packer').startup(function(use)
    -- Packer puede autogestionarse a sí mismo
    use 'wbthomason/packer.nvim'

    -- Plugins adicionales
    use 'beauwilliams/statusline.lua'            -- Statusline personalizado
    use 'nvim-lualine/lualine.nvim'              -- Otra opción de statusline
    use 'nvim-tree/nvim-web-devicons'            -- Iconos para statusline, nvim-tree, etc.
    use 'nvim-tree/nvim-tree.lua'                -- Explorador de archivos
    use "lukas-reineke/indent-blankline.nvim"    -- Muestra líneas de indentación

    -- Dashboard para Neovim
    use {
        'nvimdev/dashboard-nvim',
        event = 'VimEnter',                      -- Carga el plugin al entrar en Neovim
        config = function()
            require('dashboard').setup {
                -- Configuración del dashboard
                theme = 'hyper',                 -- Ejemplo de configuración
                disable_at_vimenter = false,     -- No deshabilitar al iniciar
            }
        end,
    }
end)

