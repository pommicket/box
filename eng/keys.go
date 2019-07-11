/*
This file is part of Box.

Box is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Box is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Box.  If not, see <https://www.gnu.org/licenses/>.
*/

package eng

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	KEY_UNKNOWN            = 0
	KEY_BACKSPACE          = 8
	KEY_TAB                = 9
	KEY_RETURN             = 13
	KEY_ESCAPE             = 27
	KEY_SPACE              = 32
	KEY_EXCLAIM            = 33
	KEY_QUOTEDBL           = 34
	KEY_HASH               = 35
	KEY_DOLLAR             = 36
	KEY_PERCENT            = 37
	KEY_AMPERSAND          = 38
	KEY_QUOTE              = 39
	KEY_LEFTPAREN          = 40
	KEY_RIGHTPAREN         = 41
	KEY_ASTERISK           = 42
	KEY_PLUS               = 43
	KEY_COMMA              = 44
	KEY_MINUS              = 45
	KEY_PERIOD             = 46
	KEY_SLASH              = 47
	KEY_0                  = 48
	KEY_1                  = 49
	KEY_2                  = 50
	KEY_3                  = 51
	KEY_4                  = 52
	KEY_5                  = 53
	KEY_6                  = 54
	KEY_7                  = 55
	KEY_8                  = 56
	KEY_9                  = 57
	KEY_COLON              = 58
	KEY_SEMICOLON          = 59
	KEY_LESS               = 60
	KEY_EQUALS             = 61
	KEY_GREATER            = 62
	KEY_QUESTION           = 63
	KEY_AT                 = 64
	KEY_LEFTBRACKET        = 91
	KEY_BACKSLASH          = 92
	KEY_RIGHTBRACKET       = 93
	KEY_CARET              = 94
	KEY_UNDERSCORE         = 95
	KEY_BACKQUOTE          = 96
	KEY_a                  = 97
	KEY_b                  = 98
	KEY_c                  = 99
	KEY_d                  = 100
	KEY_e                  = 101
	KEY_f                  = 102
	KEY_g                  = 103
	KEY_h                  = 104
	KEY_i                  = 105
	KEY_j                  = 106
	KEY_k                  = 107
	KEY_l                  = 108
	KEY_m                  = 109
	KEY_n                  = 110
	KEY_o                  = 111
	KEY_p                  = 112
	KEY_q                  = 113
	KEY_r                  = 114
	KEY_s                  = 115
	KEY_t                  = 116
	KEY_u                  = 117
	KEY_v                  = 118
	KEY_w                  = 119
	KEY_x                  = 120
	KEY_y                  = 121
	KEY_z                  = 122
	KEY_DELETE             = 127
	KEY_CAPSLOCK           = 569
	KEY_F1                 = 570
	KEY_F2                 = 571
	KEY_F3                 = 572
	KEY_F4                 = 573
	KEY_F5                 = 574
	KEY_F6                 = 575
	KEY_F7                 = 576
	KEY_F8                 = 577
	KEY_F9                 = 578
	KEY_F10                = 579
	KEY_F11                = 580
	KEY_F12                = 581
	KEY_PRINTSCREEN        = 582
	KEY_SCROLLLOCK         = 583
	KEY_PAUSE              = 584
	KEY_INSERT             = 585
	KEY_HOME               = 586
	KEY_PAGEUP             = 587
	KEY_END                = 589
	KEY_PAGEDOWN           = 590
	KEY_RIGHT              = 591
	KEY_LEFT               = 592
	KEY_DOWN               = 593
	KEY_UP                 = 594
	KEY_NUMLOCKCLEAR       = 595
	KEY_KP_DIVIDE          = 596
	KEY_KP_MULTIPLY        = 597
	KEY_KP_MINUS           = 598
	KEY_KP_PLUS            = 599
	KEY_KP_ENTER           = 600
	KEY_KP_1               = 601
	KEY_KP_2               = 602
	KEY_KP_3               = 603
	KEY_KP_4               = 604
	KEY_KP_5               = 605
	KEY_KP_6               = 606
	KEY_KP_7               = 607
	KEY_KP_8               = 608
	KEY_KP_9               = 609
	KEY_KP_0               = 610
	KEY_KP_PERIOD          = 611
	KEY_APPLICATION        = 613
	KEY_POWER              = 614
	KEY_KP_EQUALS          = 615
	KEY_F13                = 616
	KEY_F14                = 617
	KEY_F15                = 618
	KEY_F16                = 619
	KEY_F17                = 620
	KEY_F18                = 621
	KEY_F19                = 622
	KEY_F20                = 623
	KEY_F21                = 624
	KEY_F22                = 625
	KEY_F23                = 626
	KEY_F24                = 627
	KEY_EXECUTE            = 628
	KEY_HELP               = 629
	KEY_MENU               = 630
	KEY_SELECT             = 631
	KEY_STOP               = 632
	KEY_AGAIN              = 633
	KEY_UNDO               = 634
	KEY_CUT                = 635
	KEY_COPY               = 636
	KEY_PASTE              = 637
	KEY_FIND               = 638
	KEY_MUTE               = 639
	KEY_VOLUMEUP           = 640
	KEY_VOLUMEDOWN         = 641
	KEY_KP_COMMA           = 645
	KEY_KP_EQUALSAS400     = 646
	KEY_ALTERASE           = 665
	KEY_SYSREQ             = 666
	KEY_CANCEL             = 667
	KEY_CLEAR              = 668
	KEY_PRIOR              = 669
	KEY_RETURN2            = 670
	KEY_SEPARATOR          = 671
	KEY_OUT                = 672
	KEY_OPER               = 673
	KEY_CLEARAGAIN         = 674
	KEY_CRSEL              = 675
	KEY_EXSEL              = 676
	KEY_KP_00              = 688
	KEY_KP_000             = 689
	KEY_THOUSANDSSEPARATOR = 690
	KEY_DECIMALSEPARATOR   = 691
	KEY_CURRENCYUNIT       = 692
	KEY_CURRENCYSUBUNIT    = 693
	KEY_KP_LEFTPAREN       = 694
	KEY_KP_RIGHTPAREN      = 695
	KEY_KP_LEFTBRACE       = 696
	KEY_KP_RIGHTBRACE      = 697
	KEY_KP_TAB             = 698
	KEY_KP_BACKSPACE       = 699
	KEY_KP_A               = 700
	KEY_KP_B               = 701
	KEY_KP_C               = 702
	KEY_KP_D               = 703
	KEY_KP_E               = 704
	KEY_KP_F               = 705
	KEY_KP_XOR             = 706
	KEY_KP_POWER           = 707
	KEY_KP_PERCENT         = 708
	KEY_KP_LESS            = 709
	KEY_KP_GREATER         = 710
	KEY_KP_AMPERSAND       = 711
	KEY_KP_DBLAMPERSAND    = 712
	KEY_KP_VERTICALBAR     = 713
	KEY_KP_DBLVERTICALBAR  = 714
	KEY_KP_COLON           = 715
	KEY_KP_HASH            = 716
	KEY_KP_SPACE           = 717
	KEY_KP_AT              = 718
	KEY_KP_EXCLAM          = 719
	KEY_KP_MEMSTORE        = 720
	KEY_KP_MEMRECALL       = 721
	KEY_KP_MEMCLEAR        = 722
	KEY_KP_MEMADD          = 723
	KEY_KP_MEMSUBTRACT     = 724
	KEY_KP_MEMMULTIPLY     = 725
	KEY_KP_MEMDIVIDE       = 726
	KEY_KP_PLUSMINUS       = 727
	KEY_KP_CLEAR           = 728
	KEY_KP_CLEARENTRY      = 729
	KEY_KP_BINARY          = 730
	KEY_KP_OCTAL           = 731
	KEY_KP_DECIMAL         = 732
	KEY_KP_HEXADECIMAL     = 733
	KEY_LCTRL              = 736
	KEY_LSHIFT             = 737
	KEY_LALT               = 738
	KEY_LGUI               = 739
	KEY_RCTRL              = 740
	KEY_RSHIFT             = 741
	KEY_RALT               = 742
	KEY_RGUI               = 743
	KEY_MODE               = 769
	KEY_AUDIONEXT          = 770
	KEY_AUDIOPREV          = 771
	KEY_AUDIOSTOP          = 772
	KEY_AUDIOPLAY          = 773
	KEY_AUDIOMUTE          = 774
	KEY_MEDIASELECT        = 775
	KEY_WWW                = 776
	KEY_MAIL               = 777
	KEY_CALCULATOR         = 778
	KEY_COMPUTER           = 779
	KEY_AC_SEARCH          = 780
	KEY_AC_HOME            = 781
	KEY_AC_BACK            = 782
	KEY_AC_FORWARD         = 783
	KEY_AC_STOP            = 784
	KEY_AC_REFRESH         = 785
	KEY_AC_BOOKMARKS       = 786
	KEY_BRIGHTNESSDOWN     = 787
	KEY_BRIGHTNESSUP       = 788
	KEY_DISPLAYSWITCH      = 789
	KEY_KBDILLUMTOGGLE     = 790
	KEY_KBDILLUMDOWN       = 791
	KEY_KBDILLUMUP         = 792
	KEY_EJECT              = 793
	KEY_SLEEP              = 794
)

// Converts an SDL keycode to an eng keycode
func fromSdlKeycode(keycode sdl.Keycode) int {
	// Make keycode numbers smaller (why? why not.)
	code := int64(keycode)
	if (code & 0x40000000) != 0 {
		code &= ^0x400000000
		code |= 0x200
	}
	return int(code)
}

// Converts an eng keycode to an SDL keycode
func toSdlKeycode(keycode int) sdl.Keycode {
	if (keycode & 0x200) != 0 {
		keycode &= ^0x200
		keycode |= 0x40000000
	}
	return sdl.Keycode(keycode)
}
