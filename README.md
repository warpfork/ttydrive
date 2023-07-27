ttydrive
========

It's a little hack in golang which yeets "keystrokes" into a terminal device.

It's fake typing, in other words.

You can use this to automate an "interactive" demo, for example.

It has limitations.  Namely, you can't really wait for anything or read out any feedback.
Those just aren't capabilities that the TTY ABI itself really offers.
You can hack around it with other scripts however you dare; or, just use lots of `sleep`.  good luck.

Provided very much without waranty.

License: Apache v2 or MIT, at your option.
