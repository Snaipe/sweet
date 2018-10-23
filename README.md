# Sweet

A [sway][sway]-to-[i3][i3] IPC compatibility layer.

```
go get -u snai.pe/sweet
```

# Why

Sway provides a value for I3SOCK, but the json output it sends can vary
slightly from stock i3. This sometimes break programs consuming this API,
because they expect some fields to be always present, even if they don't
quite make sense in a pure wayland sense.

Sweet solves this problem by creating an IPC socket at $SWAYSOCK.i3, which
tries to mimic as much as it can i3's output.

# Usage

The simplest way to set up sweet is to add to your sway config:

```
exec sweet
```

and then set `I3SOCK="$SWAYSOCK.i3"` in your environment.

# Changes

* The class and the instance of a wayland window will always be the app ID of the view.
* The `app_id` field is removed from view objects
* View objects of wayland windows provide the `window_properties` map.

[sway]: https://swaywm.org/
[i3]: https://i3wm.org/
