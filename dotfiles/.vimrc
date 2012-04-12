set nocompatible
filetype off

set rtp+=~/.vim/bundle/vundle/
call vundle#rc()

" insert Bundle plugins here

syntax on
filetype plugin indent on

set scrolloff=5
set laststatus=2
set textwidth=0
set autoindent shiftwidth=2
set smartindent shiftwidth=2
set expandtab
set list
set listchars=tab:>.
set hlsearch

" encodings 
set encoding=utf8
set fileencodings=iso-2022-jp,sjis,euc-jp,utf8

" keymap 
nnoremap <silent> t :<C-u>tabnew<CR>:tabmove<CR>
inoremap jj <Esc>
inoremap kk <Esc>
" not use arrow keys
noremap <Up> <nop>
noremap <Down> <nop>
noremap <Left> <nop>
noremap <Right> <nop>


:au BufReadPost * if line("'\"") > 1 && line("'\"") <= line("$") | exe "normal! g`\"" | endif
:au Syntax go source $NISE_GO_HOME/misc/vim/syntax/go.vim
autocmd BufNewFile,BufRead *.go set syntax=go

"" plugins 

" plugin gist.vim
let g:github_user='nise-nabe'
let g:github_token='nise-token'


" plugin vimfiler
let g:vimfiler_safe_mode_by_default = 0
let g:vimfiler_as_default_explorer = 1

" "plugin syntastic
let g:syntastic_check_on_open = 1

" plugin ack.vim
let g:ackprg="ack -H --nocolor --nogroup --column"

" plugin calendar.vim
let g:calendar_wruler = "日 月 火 水 木 金 土"
let g:calendar_diary = "~/.diary"
let calendar_action = "QFixHowmCalendarDiary"
let calendar_sign = "QFixHowmCalendarSign"

" plugin qfixhowm
let howm_dir = "~/.howm"
let howm_fileencoding = "utf-8"
let howm_fileformat = "unix"
let QFixHowm_HowmMode = 0
let QFixHowm_Title = '#'
let QFixHowm_UserFileType = 'markdown'
let QFixHowm_UserFileExt = 'mkd'
let howm_filename = '%Y/%m/%T-%m-%d-%H%M%S.mkd'
