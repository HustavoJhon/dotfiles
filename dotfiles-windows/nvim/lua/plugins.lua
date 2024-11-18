require('packer').startup(function()
	-- Packer can manager itself
	use 'beauwilliams/statusline.lua'
	use 'nvim-tree/nvim-web-devicons'
	use 'nvim-tree/nvim-tree.lua'
	use "lukas-reineke/indent-blankline.nvim"
    use {
  		'nvimdev/dashboard-nvim',
  		event = 'VimEnter',
  		config = function()
    		require('dashboard').setup {
      	-- config
    	}
  	end,
	}
end)


