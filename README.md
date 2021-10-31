# i3window

Small utility for use with [polybar][] and [i3][], which outputs the name of
the focused window on the given X11 output. The intent is to create a widget
for [polybar], with the behavior that I prefer.

## Usage

    go install github.com/jdelkins/i3window
    i3window <display-name>

`<display-name>` should be one of the display names returned with `polybar
--list-all-monitors`.  This manual running approach is fine for testing, but
the program is meant to be used with [polybar], so see below for how to
configure and run from there.

## How it Works

Whenever the focused window changes in the provided window, either through
opening, closing, or focusing a window, the program changing the window title,
or by changing workspaces on that monitor, the program will output the focused
X11 window name to stdout. When [polybar] is configured with a module that runs
this program as a custom script, it will show this output on the bar.

## [Polybar][] Configuration

Example configuration follows. The important items are `type = custom/script`,
`exec = ...` and `tail = true`.

    [module/i3-window]
    type = custom/script
    label = %output%
    tail = true
    exec = ~/go/bin/i3window $MONITOR

## Acknowledgements

Uses the indespensible [i3ipc](https://github.com/mdirkse/i3ipc) package to
interface with i3. The [polybar] interface is handled by [polybar] itself,
which simply reads from `stdout`.

## License

_i3window_, a program to show window title information from [i3] in [polybar].
Copyright (c) 2021 Joel D. Elkins <joel@elkins.co>.

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
[GNU General Public License][GPL] for more details.

You should have received a copy of the [GNU General Public License][GPL]
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

[i3]: https://i3wm.org/
[polybar]: https://github.com/polybar/polybar
[GPL]: https://www.gnu.org/licenses/old-licenses/gpl-2.0.html
