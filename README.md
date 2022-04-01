# gibson-cli
Golang CLI tool for Gibson framework: package manager, project manager, godot cli parser, class resolver, etc...

## Modules
- [x] [Package Manager](#package-manager)  
- [ ] [Project Manager]()  
- [ ] [CLI Parser]()  
- [ ] [Class Resolver]()  

## Installation
### with GO
If you have [go already installed](https://go.dev/doc/install) on your machine, clone this repository, and then:
```bash
cd gibson-cli
go install
```
and gibson will be ready to serve you!

### without GO
Download binaries at your choice from the [latest release](https://github.com/gibsongd/gibson-cli/releases).
Unzip gibson binaries wherever you want and
- on linux, move `gibson` to `/usr/bin`
- on windows, add gibson folder to %PATH% environment variable

### Package Manager
Gibson package manager lets you install addons in your Godot projects.
Gibson handles assets in a very similar (yet liter) way to `npm`.
Addons are installed only from one trusted source, which is the official [AssetLib](https://godotengine.org/asset-library/asset) using their public APIs.
Once installed, addons will be cached in your %APP_DATA% folder, so all your Godot projects will be able to share addons just executing one command!

#### examples

*search for an addon*
```bash
> gibson pm -search "time"

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
> gibson pm -install "pycbouh/Time Tracker"

[gibson-cli] ✓
[gibson-cli] ✓   `pycbouh/Time Tracker` fetched!
[gibson-cli] ✓   `662` info retrieved!
[gibson-cli] ✓   `pycbouh/Time Tracker` downloaded!
[gibson-cli] ✓   `pycbouh/Time Tracker` installed successfully!
```

*install an addon by `id`*
```bash
> gibson pm -install 662

[gibson-cli] ✓
[gibson-cli] ✓   `662` info retrieved!
[gibson-cli] ✓   `pycbouh/Time Tracker` downloaded!
[gibson-cli] ✓   `pycbouh/Time Tracker` installed successfully!
```

When an addon is installed, it will create a `gibson.json` in the current folder.
It is used by *gibson* in order to handle addons.
If an addon is already cached, it will just be unpacked in the current project.

*install an addon by `author/title` (cached)*
```bash
> gibson pm -install "pycbouh/Time Tracker"

[gibson-cli] ✓   `pycbouh/Time Tracker` found in cache
[gibson-cli] ✓   `pycbouh/Time Tracker` installed successfully!
```

*Want to clear the cache and force install?*
```bash
> gibson pm -install "pycbouh/Time Tracker" -clear

[gibson-cli] ✓   `pycbouh/Time Tracker` cache cleared
[gibson-cli] ✓   `pycbouh/Time Tracker` fetched!
[gibson-cli] ✓   `662` info retrieved!
[gibson-cli] ✓   `pycbouh/Time Tracker` downloaded!
[gibson-cli] ✓   `pycbouh/Time Tracker` installed successfully!
```

*uninstall an addon*
```bash
> gibson pm -uninstall "pycbouh/Time Tracker"

[gibson-cli] ✓   `time-tracker` removed
[gibson-cli] ✓   `pycbouh/Time Tracker` uninstalled
```

*install all the addons listed in the `gibson.json` file*
```bash
> gibson pm -install .
```
This is useful if you want to share your project and leave it addons-independent or want to make a lighter project.

*list all addons currently installed (using gibson)*
```bash
> gibson pm -list
```
