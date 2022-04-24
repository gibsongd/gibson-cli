# gibson-cli
Golang CLI tool for Gibson framework: package manager, project manager, godot cli parser, class resolver, etc...

## Modules
- [x] [Package Manager](#package-manager)  
- [ ] [Project Manager]()  
- [ ] [CLI Parser]()  
- [ ] [Class Resolver]()  

## Installation
### with GO
If you have [go already installed](https://go.dev/doc/install) on your machine:
```bash
go install https://github.com/gibsongd/gibson-cli
```
and gibson will be ready to serve you!

### without GO
Download binaries at your choice from the [latest release](https://github.com/gibsongd/gibson-cli/releases).
Unzip gibson binaries wherever you want and
- on linux, move `gibson` binaries to `/usr/bin`
- on windows, add `gibson.exe` folder to `%PATH%` environment variable

### Package Manager
Gibson package manager lets you install addons in your Godot projects.
Gibson handles assets in a very similar (yet liter) way to `npm`.
Addons are installed only from one trusted source, which is the official [AssetLib](https://godotengine.org/asset-library/asset) using their public APIs.
Once installed, addons will be cached in your `$HOME/.cache` (*unix*) or `%LocalAppData%` (*windows*) folder, so all your Godot projects will be able to share addons just executing one command!

#### examples

*search for an addon*
```bash
> gibson pm search time
or
> gibson pm search "time"

[id]    [user]          [title]
1275    Blaron          Timer Counter
1157    Aendryr         Date Time Addon
1134    GianptDev       EditorTimer
1127    RipeX           Simple Project Timer
236     thomazthz       Godot Wakatime
721     VolodyaKEK      Single File Runtime Console
702     Eminnet         EminMultiMesh
662     pycbouh         Time Tracker
342     Cevantime       Water2D Node
```

*install an addon by `author/title`*
```bash
> gibson package_manager install "fenix/Godot Engine JWT"
or
> gibson pm i "fenix/Godot Engine JWT"

[gibson-cli] ✓   `fenix/Godot Engine JWT` found!
[gibson-cli] ✓   `fenix/Godot Engine JWT` info retrieved!
[gibson-cli] ✓   `fenix/Godot Engine JWT` downloaded!
[gibson-cli] ✓   `fenix/Godot Engine JWT` installed successfully!

```

*install an addon by `id`*
```bash
> gibson pm i 1104

[gibson-cli] ✓   `fenix/Godot Engine JWT` info retrieved!
[gibson-cli] ✓   `fenix/Godot Engine JWT` downloaded!
[gibson-cli] ✓   `fenix/Godot Engine JWT` installed successfully!
```

When an addon is installed, it will create a `gibson.json` in the current folder.
It is used by *gibson* in order to handle addons.
If an addon is already cached, it will just be unpacked in the current project.

*install an addon by `author/title` (cached)*
```bash
> gibson pm i "fenix/Godot Engine JWT"

[gibson-cli] ✓   `fenix/Godot Engine JWT` found in cache
[gibson-cli] ✓   `fenix/Godot Engine JWT` installed successfully!
```

*Want to force install?*
```bash
> gibson pm i "fenix/Godot Engine JWT" --force || -f
```

*uninstall an addon*
```bash
> gibson pm uninstall "fenix/Godot Engine JWT"

[gibson-cli] ✓   `fenix/Godot Engine JWT` removed
[gibson-cli] ✓   `fenix/Godot Engine JWT` uninstalled
```

*uninstall an addon and clear all the cached versions *
```bash
> gibson pm uninstall "fenix/Godot Engine JWT" --clear || -c

[gibson-cli] ✓   `fenix/Godot Engine JWT` cache cleared
[gibson-cli] ✓   `fenix/Godot Engine JWT` removed
[gibson-cli] ✓   `fenix/Godot Engine JWT` uninstalled
```

*install all the addons listed in the `gibson.json` file*
```bash
> gibson pm i
```
This is useful if you want to share your project and leave it addons-independent or want to make a lighter project.

*list all addons currently installed (using gibson)*
```bash
> gibson pm list

assets/
├─ dialogue_editor/     1200    (VP-GAMES/Dialogue Editor (G4))
├─ draw3d/              1301    (nyxkn/Draw3D)
├─ jwt/                 1104    (fenix/Godot Engine JWT)
├─ stream_comment/      1300    (velop/Stream Comment)
└─ telegram-bot-api/    1072    (fenix/Telegram Bot API)
```
