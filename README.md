# gibson-cli
Golang CLI tool for Gibson framework: package manager, project manager, godot cli parser, class resolver, etc...

## Modules

### Package Manager
Gibson package manager lets you install addons in your Godot projects.
Gibson handles assets in a very similar (yet liter) way to `npm`.
Addons are installed only from one trusted source, which is the official [AssetLib](https://godotengine.org/asset-library/asset) using their public APIs.
Once installed, addons will be cached in your %APP_DATA% folder, so all your Godot projects will be able to share addons just executing one command!

**PoC**

*install an addon by `author/title`*

*install an addon by `id`*

If an addon is already cached, it will just be unzipped in the current project.
