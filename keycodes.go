package keycodes

import "github.com/veandco/go-sdl2/sdl"

type Keycode int

const (
	K_UNKNOWN Keycode = iota
	K_RETURN
	K_ESCAPE
	K_BACKSPACE
	K_TAB
	K_SPACE
	K_EXCLAIM
	K_QUOTEDBL
	K_HASH
	K_PERCENT
	K_DOLLAR
	K_AMPERSAND
	K_QUOTE
	K_LEFTPAREN
	K_RIGHTPAREN
	K_ASTERISK
	K_PLUS
	K_COMMA
	K_MINUS
	K_PERIOD
	K_SLASH
	K_0
	K_1
	K_2
	K_3
	K_4
	K_5
	K_6
	K_7
	K_8
	K_9
	K_COLON
	K_SEMICOLON
	K_LESS
	K_EQUALS
	K_GREATER
	K_QUESTION
	K_AT
	K_LEFTBRACKET
	K_BACKSLASH
	K_RIGHTBRACKET
	K_CARET
	K_UNDERSCORE
	K_BACKQUOTE
	K_a
	K_b
	K_c
	K_d
	K_e
	K_f
	K_g
	K_h
	K_i
	K_j
	K_k
	K_l
	K_m
	K_n
	K_o
	K_p
	K_q
	K_r
	K_s
	K_t
	K_u
	K_v
	K_w
	K_x
	K_y
	K_z
	K_CAPSLOCK
	K_F1
	K_F2
	K_F3
	K_F4
	K_F5
	K_F6
	K_F7
	K_F8
	K_F9
	K_F10
	K_F11
	K_F12
	K_PRINTSCREEN
	K_SCROLLLOCK
	K_PAUSE
	K_INSERT
	K_HOME
	K_PAGEUP
	K_DELETE
	K_END
	K_PAGEDOWN
	K_RIGHT
	K_LEFT
	K_DOWN
	K_UP
	K_NUMLOCKCLEAR
	K_KP_DIVIDE
	K_KP_MULTIPLY
	K_KP_MINUS
	K_KP_PLUS
	K_KP_ENTER
	K_KP_1
	K_KP_2
	K_KP_3
	K_KP_4
	K_KP_5
	K_KP_6
	K_KP_7
	K_KP_8
	K_KP_9
	K_KP_0
	K_KP_PERIOD
	K_APPLICATION
	K_POWER
	K_KP_EQUALS
	K_F13
	K_F14
	K_F15
	K_F16
	K_F17
	K_F18
	K_F19
	K_F20
	K_F21
	K_F22
	K_F23
	K_F24
	K_EXECUTE
	K_HELP
	K_MENU
	K_SELECT
	K_STOP
	K_AGAIN
	K_UNDO
	K_CUT
	K_COPY
	K_PASTE
	K_FIND
	K_MUTE
	K_VOLUMEUP
	K_VOLUMEDOWN
	K_KP_COMMA
	K_KP_EQUALSAS400
	K_ALTERASE
	K_SYSREQ
	K_CANCEL
	K_CLEAR
	K_PRIOR
	K_RETURN2
	K_SEPARATOR
	K_OUT
	K_OPER
	K_CLEARAGAIN
	K_CRSEL
	K_EXSEL
	K_KP_00
	K_KP_000
	K_THOUSANDSSEPARATOR
	K_DECIMALSEPARATOR
	K_CURRENCYUNIT
	K_CURRENCYSUBUNIT
	K_KP_LEFTPAREN
	K_KP_RIGHTPAREN
	K_KP_LEFTBRACE
	K_KP_RIGHTBRACE
	K_KP_TAB
	K_KP_BACKSPACE
	K_KP_A
	K_KP_B
	K_KP_C
	K_KP_D
	K_KP_E
	K_KP_F
	K_KP_XOR
	K_KP_POWER
	K_KP_PERCENT
	K_KP_LESS
	K_KP_GREATER
	K_KP_AMPERSAND
	K_KP_DBLAMPERSAND
	K_KP_VERTICALBAR
	K_KP_DBLVERTICALBAR
	K_KP_COLON
	K_KP_HASH
	K_KP_SPACE
	K_KP_AT
	K_KP_EXCLAM
	K_KP_MEMSTORE
	K_KP_MEMRECALL
	K_KP_MEMCLEAR
	K_KP_MEMADD
	K_KP_MEMSUBTRACT
	K_KP_MEMMULTIPLY
	K_KP_MEMDIVIDE
	K_KP_PLUSMINUS
	K_KP_CLEAR
	K_KP_CLEARENTRY
	K_KP_BINARY
	K_KP_OCTAL
	K_KP_DECIMAL
	K_KP_HEXADECIMAL
	K_LCTRL
	K_LSHIFT
	K_LALT
	K_LGUI
	K_RCTRL
	K_RSHIFT
	K_RALT
	K_RGUI
	K_MODE
	K_AUDIONEXT
	K_AUDIOPREV
	K_AUDIOSTOP
	K_AUDIOPLAY
	K_AUDIOMUTE
	K_MEDIASELECT
	K_WWW
	K_MAIL
	K_CALCULATOR
	K_COMPUTER
	K_AC_SEARCH
	K_AC_HOME
	K_AC_BACK
	K_AC_FORWARD
	K_AC_STOP
	K_AC_REFRESH
	K_AC_BOOKMARKS
	K_BRIGHTNESSDOWN
	K_BRIGHTNESSUP
	K_DISPLAYSWITCH
	K_KBDILLUMTOGGLE
	K_KBDILLUMDOWN
	K_KBDILLUMUP
	K_EJECT
	K_SLEEP
)

func FromSDL(sym sdl.Keycode) Keycode {
	switch sym {
	case sdl.K_UNKNOWN:
		return K_UNKNOWN
	case sdl.K_RETURN:
		return K_RETURN
	case sdl.K_ESCAPE:
		return K_ESCAPE
	case sdl.K_BACKSPACE:
		return K_BACKSPACE
	case sdl.K_TAB:
		return K_TAB
	case sdl.K_SPACE:
		return K_SPACE
	case sdl.K_EXCLAIM:
		return K_EXCLAIM
	case sdl.K_QUOTEDBL:
		return K_QUOTEDBL
	case sdl.K_HASH:
		return K_HASH
	case sdl.K_PERCENT:
		return K_PERCENT
	case sdl.K_DOLLAR:
		return K_DOLLAR
	case sdl.K_AMPERSAND:
		return K_AMPERSAND
	case sdl.K_QUOTE:
		return K_QUOTE
	case sdl.K_LEFTPAREN:
		return K_LEFTPAREN
	case sdl.K_RIGHTPAREN:
		return K_RIGHTPAREN
	case sdl.K_ASTERISK:
		return K_ASTERISK
	case sdl.K_PLUS:
		return K_PLUS
	case sdl.K_COMMA:
		return K_COMMA
	case sdl.K_MINUS:
		return K_MINUS
	case sdl.K_PERIOD:
		return K_PERIOD
	case sdl.K_SLASH:
		return K_SLASH
	case sdl.K_0:
		return K_0
	case sdl.K_1:
		return K_1
	case sdl.K_2:
		return K_2
	case sdl.K_3:
		return K_3
	case sdl.K_4:
		return K_4
	case sdl.K_5:
		return K_5
	case sdl.K_6:
		return K_6
	case sdl.K_7:
		return K_7
	case sdl.K_8:
		return K_8
	case sdl.K_9:
		return K_9
	case sdl.K_COLON:
		return K_COLON
	case sdl.K_SEMICOLON:
		return K_SEMICOLON
	case sdl.K_LESS:
		return K_LESS
	case sdl.K_EQUALS:
		return K_EQUALS
	case sdl.K_GREATER:
		return K_GREATER
	case sdl.K_QUESTION:
		return K_QUESTION
	case sdl.K_AT:
		return K_AT
	case sdl.K_LEFTBRACKET:
		return K_LEFTBRACKET
	case sdl.K_BACKSLASH:
		return K_BACKSLASH
	case sdl.K_RIGHTBRACKET:
		return K_RIGHTBRACKET
	case sdl.K_CARET:
		return K_CARET
	case sdl.K_UNDERSCORE:
		return K_UNDERSCORE
	case sdl.K_BACKQUOTE:
		return K_BACKQUOTE
	case sdl.K_a:
		return K_a
	case sdl.K_b:
		return K_b
	case sdl.K_c:
		return K_c
	case sdl.K_d:
		return K_d
	case sdl.K_e:
		return K_e
	case sdl.K_f:
		return K_f
	case sdl.K_g:
		return K_g
	case sdl.K_h:
		return K_h
	case sdl.K_i:
		return K_i
	case sdl.K_j:
		return K_j
	case sdl.K_k:
		return K_k
	case sdl.K_l:
		return K_l
	case sdl.K_m:
		return K_m
	case sdl.K_n:
		return K_n
	case sdl.K_o:
		return K_o
	case sdl.K_p:
		return K_p
	case sdl.K_q:
		return K_q
	case sdl.K_r:
		return K_r
	case sdl.K_s:
		return K_s
	case sdl.K_t:
		return K_t
	case sdl.K_u:
		return K_u
	case sdl.K_v:
		return K_v
	case sdl.K_w:
		return K_w
	case sdl.K_x:
		return K_x
	case sdl.K_y:
		return K_y
	case sdl.K_z:
		return K_z
	case sdl.K_CAPSLOCK:
		return K_CAPSLOCK
	case sdl.K_F1:
		return K_F1
	case sdl.K_F2:
		return K_F2
	case sdl.K_F3:
		return K_F3
	case sdl.K_F4:
		return K_F4
	case sdl.K_F5:
		return K_F5
	case sdl.K_F6:
		return K_F6
	case sdl.K_F7:
		return K_F7
	case sdl.K_F8:
		return K_F8
	case sdl.K_F9:
		return K_F9
	case sdl.K_F10:
		return K_F10
	case sdl.K_F11:
		return K_F11
	case sdl.K_F12:
		return K_F12
	case sdl.K_PRINTSCREEN:
		return K_PRINTSCREEN
	case sdl.K_SCROLLLOCK:
		return K_SCROLLLOCK
	case sdl.K_PAUSE:
		return K_PAUSE
	case sdl.K_INSERT:
		return K_INSERT
	case sdl.K_HOME:
		return K_HOME
	case sdl.K_PAGEUP:
		return K_PAGEUP
	case sdl.K_DELETE:
		return K_DELETE
	case sdl.K_END:
		return K_END
	case sdl.K_PAGEDOWN:
		return K_PAGEDOWN
	case sdl.K_RIGHT:
		return K_RIGHT
	case sdl.K_LEFT:
		return K_LEFT
	case sdl.K_DOWN:
		return K_DOWN
	case sdl.K_UP:
		return K_UP
	case sdl.K_NUMLOCKCLEAR:
		return K_NUMLOCKCLEAR
	case sdl.K_KP_DIVIDE:
		return K_KP_DIVIDE
	case sdl.K_KP_MULTIPLY:
		return K_KP_MULTIPLY
	case sdl.K_KP_MINUS:
		return K_KP_MINUS
	case sdl.K_KP_PLUS:
		return K_KP_PLUS
	case sdl.K_KP_ENTER:
		return K_KP_ENTER
	case sdl.K_KP_1:
		return K_KP_1
	case sdl.K_KP_2:
		return K_KP_2
	case sdl.K_KP_3:
		return K_KP_3
	case sdl.K_KP_4:
		return K_KP_4
	case sdl.K_KP_5:
		return K_KP_5
	case sdl.K_KP_6:
		return K_KP_6
	case sdl.K_KP_7:
		return K_KP_7
	case sdl.K_KP_8:
		return K_KP_8
	case sdl.K_KP_9:
		return K_KP_9
	case sdl.K_KP_0:
		return K_KP_0
	case sdl.K_KP_PERIOD:
		return K_KP_PERIOD
	case sdl.K_APPLICATION:
		return K_APPLICATION
	case sdl.K_POWER:
		return K_POWER
	case sdl.K_KP_EQUALS:
		return K_KP_EQUALS
	case sdl.K_F13:
		return K_F13
	case sdl.K_F14:
		return K_F14
	case sdl.K_F15:
		return K_F15
	case sdl.K_F16:
		return K_F16
	case sdl.K_F17:
		return K_F17
	case sdl.K_F18:
		return K_F18
	case sdl.K_F19:
		return K_F19
	case sdl.K_F20:
		return K_F20
	case sdl.K_F21:
		return K_F21
	case sdl.K_F22:
		return K_F22
	case sdl.K_F23:
		return K_F23
	case sdl.K_F24:
		return K_F24
	case sdl.K_EXECUTE:
		return K_EXECUTE
	case sdl.K_HELP:
		return K_HELP
	case sdl.K_MENU:
		return K_MENU
	case sdl.K_SELECT:
		return K_SELECT
	case sdl.K_STOP:
		return K_STOP
	case sdl.K_AGAIN:
		return K_AGAIN
	case sdl.K_UNDO:
		return K_UNDO
	case sdl.K_CUT:
		return K_CUT
	case sdl.K_COPY:
		return K_COPY
	case sdl.K_PASTE:
		return K_PASTE
	case sdl.K_FIND:
		return K_FIND
	case sdl.K_MUTE:
		return K_MUTE
	case sdl.K_VOLUMEUP:
		return K_VOLUMEUP
	case sdl.K_VOLUMEDOWN:
		return K_VOLUMEDOWN
	case sdl.K_KP_COMMA:
		return K_KP_COMMA
	case sdl.K_KP_EQUALSAS400:
		return K_KP_EQUALSAS400
	case sdl.K_ALTERASE:
		return K_ALTERASE
	case sdl.K_SYSREQ:
		return K_SYSREQ
	case sdl.K_CANCEL:
		return K_CANCEL
	case sdl.K_CLEAR:
		return K_CLEAR
	case sdl.K_PRIOR:
		return K_PRIOR
	case sdl.K_RETURN2:
		return K_RETURN2
	case sdl.K_SEPARATOR:
		return K_SEPARATOR
	case sdl.K_OUT:
		return K_OUT
	case sdl.K_OPER:
		return K_OPER
	case sdl.K_CLEARAGAIN:
		return K_CLEARAGAIN
	case sdl.K_CRSEL:
		return K_CRSEL
	case sdl.K_EXSEL:
		return K_EXSEL
	case sdl.K_KP_00:
		return K_KP_00
	case sdl.K_KP_000:
		return K_KP_000
	case sdl.K_THOUSANDSSEPARATOR:
		return K_THOUSANDSSEPARATOR
	case sdl.K_DECIMALSEPARATOR:
		return K_DECIMALSEPARATOR
	case sdl.K_CURRENCYUNIT:
		return K_CURRENCYUNIT
	case sdl.K_CURRENCYSUBUNIT:
		return K_CURRENCYSUBUNIT
	case sdl.K_KP_LEFTPAREN:
		return K_KP_LEFTPAREN
	case sdl.K_KP_RIGHTPAREN:
		return K_KP_RIGHTPAREN
	case sdl.K_KP_LEFTBRACE:
		return K_KP_LEFTBRACE
	case sdl.K_KP_RIGHTBRACE:
		return K_KP_RIGHTBRACE
	case sdl.K_KP_TAB:
		return K_KP_TAB
	case sdl.K_KP_BACKSPACE:
		return K_KP_BACKSPACE
	case sdl.K_KP_A:
		return K_KP_A
	case sdl.K_KP_B:
		return K_KP_B
	case sdl.K_KP_C:
		return K_KP_C
	case sdl.K_KP_D:
		return K_KP_D
	case sdl.K_KP_E:
		return K_KP_E
	case sdl.K_KP_F:
		return K_KP_F
	case sdl.K_KP_XOR:
		return K_KP_XOR
	case sdl.K_KP_POWER:
		return K_KP_POWER
	case sdl.K_KP_PERCENT:
		return K_KP_PERCENT
	case sdl.K_KP_LESS:
		return K_KP_LESS
	case sdl.K_KP_GREATER:
		return K_KP_GREATER
	case sdl.K_KP_AMPERSAND:
		return K_KP_AMPERSAND
	case sdl.K_KP_DBLAMPERSAND:
		return K_KP_DBLAMPERSAND
	case sdl.K_KP_VERTICALBAR:
		return K_KP_VERTICALBAR
	case sdl.K_KP_DBLVERTICALBAR:
		return K_KP_DBLVERTICALBAR
	case sdl.K_KP_COLON:
		return K_KP_COLON
	case sdl.K_KP_HASH:
		return K_KP_HASH
	case sdl.K_KP_SPACE:
		return K_KP_SPACE
	case sdl.K_KP_AT:
		return K_KP_AT
	case sdl.K_KP_EXCLAM:
		return K_KP_EXCLAM
	case sdl.K_KP_MEMSTORE:
		return K_KP_MEMSTORE
	case sdl.K_KP_MEMRECALL:
		return K_KP_MEMRECALL
	case sdl.K_KP_MEMCLEAR:
		return K_KP_MEMCLEAR
	case sdl.K_KP_MEMADD:
		return K_KP_MEMADD
	case sdl.K_KP_MEMSUBTRACT:
		return K_KP_MEMSUBTRACT
	case sdl.K_KP_MEMMULTIPLY:
		return K_KP_MEMMULTIPLY
	case sdl.K_KP_MEMDIVIDE:
		return K_KP_MEMDIVIDE
	case sdl.K_KP_PLUSMINUS:
		return K_KP_PLUSMINUS
	case sdl.K_KP_CLEAR:
		return K_KP_CLEAR
	case sdl.K_KP_CLEARENTRY:
		return K_KP_CLEARENTRY
	case sdl.K_KP_BINARY:
		return K_KP_BINARY
	case sdl.K_KP_OCTAL:
		return K_KP_OCTAL
	case sdl.K_KP_DECIMAL:
		return K_KP_DECIMAL
	case sdl.K_KP_HEXADECIMAL:
		return K_KP_HEXADECIMAL
	case sdl.K_LCTRL:
		return K_LCTRL
	case sdl.K_LSHIFT:
		return K_LSHIFT
	case sdl.K_LALT:
		return K_LALT
	case sdl.K_LGUI:
		return K_LGUI
	case sdl.K_RCTRL:
		return K_RCTRL
	case sdl.K_RSHIFT:
		return K_RSHIFT
	case sdl.K_RALT:
		return K_RALT
	case sdl.K_RGUI:
		return K_RGUI
	case sdl.K_MODE:
		return K_MODE
	case sdl.K_AUDIONEXT:
		return K_AUDIONEXT
	case sdl.K_AUDIOPREV:
		return K_AUDIOPREV
	case sdl.K_AUDIOSTOP:
		return K_AUDIOSTOP
	case sdl.K_AUDIOPLAY:
		return K_AUDIOPLAY
	case sdl.K_AUDIOMUTE:
		return K_AUDIOMUTE
	case sdl.K_MEDIASELECT:
		return K_MEDIASELECT
	case sdl.K_WWW:
		return K_WWW
	case sdl.K_MAIL:
		return K_MAIL
	case sdl.K_CALCULATOR:
		return K_CALCULATOR
	case sdl.K_COMPUTER:
		return K_COMPUTER
	case sdl.K_AC_SEARCH:
		return K_AC_SEARCH
	case sdl.K_AC_HOME:
		return K_AC_HOME
	case sdl.K_AC_BACK:
		return K_AC_BACK
	case sdl.K_AC_FORWARD:
		return K_AC_FORWARD
	case sdl.K_AC_STOP:
		return K_AC_STOP
	case sdl.K_AC_REFRESH:
		return K_AC_REFRESH
	case sdl.K_AC_BOOKMARKS:
		return K_AC_BOOKMARKS
	case sdl.K_BRIGHTNESSDOWN:
		return K_BRIGHTNESSDOWN
	case sdl.K_BRIGHTNESSUP:
		return K_BRIGHTNESSUP
	case sdl.K_DISPLAYSWITCH:
		return K_DISPLAYSWITCH
	case sdl.K_KBDILLUMTOGGLE:
		return K_KBDILLUMTOGGLE
	case sdl.K_KBDILLUMDOWN:
		return K_KBDILLUMDOWN
	case sdl.K_KBDILLUMUP:
		return K_KBDILLUMUP
	case sdl.K_EJECT:
		return K_EJECT
	case sdl.K_SLEEP:
		return K_SLEEP
	}
}
