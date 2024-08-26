# Manjaro Sway Edition

[![downloads](https://img.shields.io/badge/dynamic/json?color=green&label=%E2%AC%87%20%E2%88%91%20%E2%88%9E&cache=3600&query=count&url=https%3A%2F%2Fmanjaro-sway.download/count)](https://manjaro-sway.download)
[![downloads last seven days](https://img.shields.io/badge/dynamic/json?color=green&label=%E2%AC%87%20%E2%88%91%207d&cache=3600&query=sevenDays&url=https%3A%2F%2Fmanjaro-sway.download/count)](https://manjaro-sway.download)
[![downloads per week](https://img.shields.io/badge/dynamic/json?color=green&label=%E2%AC%87%20%E2%8C%80%20week&cache=3600&query=weeklyAverage&url=https%3A%2F%2Fmanjaro-sway.download/count)](https://manjaro-sway.download)

[![settings release](https://img.shields.io/github/v/release/manjaro-sway/desktop-settings)](https://github.com/Manjaro-Sway/desktop-settings/releases/latest)
[![lts](https://img.shields.io/badge/dynamic/json?label=lts&query=%24%5B%3A1%5D.packageName&url=https%3A%2F%2Fkernel-info.manjaro-sway.download%2F%3Fcategory%3Dlongterm)](https://github.com/Manjaro-Sway/manjaro-sway/releases/latest)
[![stable](https://img.shields.io/badge/dynamic/json?label=stable&query=%24%5B%3A1%5D.packageName&url=https%3A%2F%2Fkernel-info.manjaro-sway.download%2F%3Fcategory%3Dstable)](https://github.com/Manjaro-Sway/manjaro-sway/releases/latest)

[![repo build](https://github.com/manjaro-sway/packages/workflows/repo-add/badge.svg?event=repository_dispatch)](https://github.com/manjaro-sway/packages/actions)
[![iso build](https://github.com/Manjaro-Sway/manjaro-sway/actions/workflows/iso_build.yaml/badge.svg)](https://github.com/Manjaro-Sway/manjaro-sway/actions/workflows/iso_build.yaml)

[![All Contributors](https://img.shields.io/badge/dynamic/json?color=important&label=contributors&query=%24.contributors.length&url=https%3A%2F%2Fraw.githubusercontent.com%2FManjaro-Sway%2Fmanjaro-sway%2Fmain%2F.all-contributorsrc)](#contributors-)
[![Matrix](https://img.shields.io/matrix/manjaro-sway:matrix.org)](https://matrix.to/#/#manjaro-sway:matrix.org)

![image](https://user-images.githubusercontent.com/4662748/212540664-1445ae40-96d2-48b6-92e1-d9c0d350b420.png)

We are building a manjaro sway edition - with the following principles:

- use a decent cli/tui solution
- convention overrideable by configuration
- prepare opt-out
- build everything automatically

## How to install

You can find the latest images on [manjaro-sway.download](https://manjaro-sway.download/).

You can create a bootable USB stick using [Etcher](https://www.balena.io/etcher/) or a similar tool.

Check out our [FAQ](SUPPORT.md) for additional hints.

## Development

### Sources

- [iso profile](https://github.com/manjaro-sway/iso-profiles/tree/sway/community/sway)
- [desktop settings](https://github.com/manjaro-sway/desktop-settings/tree/sway/community/sway)

### How to Build

1. Check out the ISO profile
2. Run this command in the root directory

```bash
buildiso -p sway
```

### Contributing

There are lots of ways to contribute. 

- Give us a ⭐ here on github to increase our visibility
- Help collecting implementation ideas in [discussions](https://github.com/Manjaro-Sway/manjaro-sway/discussions)
- Implement ideas in our [desktop-settings](https://github.com/manjaro-sway/desktop-settings/tree/sway/community/sway) and [iso profile](https://github.com/manjaro-sway/iso-profiles/tree/sway/community/sway) and create pull requests
- Contribute to the documentation and help others in our chat
- Get in [touch](https://forum.manjaro.org/) with the broader Manjaro community.
- Use the distribution on a daily basis, find and share solutions to problems you have. 
- Get the manjaro packages before they are released to the general public by [switching to our "unstable" or "testing" branch](https://wiki.manjaro.org/index.php/Switching_Branches#Changing_to_another_branch) and report issues you face early on.

### Credits

- initial inspiration came from the [sway branch in the manjaro iso profiles repo](https://gitlab.manjaro.org/profiles-and-settings/iso-profiles/-/tree/sway)
- initially a lot of work got copied from the [manjaro sway arm overlay](https://gitlab.manjaro.org/manjaro-arm/applications/arm-profiles/-/tree/master/overlays/sway)
- the background image is [beautifully made by reddit user u/atlas-ark](https://www.reddit.com/r/wallpaper/comments/kmh680/1920x1080_all_resolutions_available_dark_light/?utm_source=share&utm_medium=web2x&context=3)
- the logo is a contribution by [André Vallestero](https://github.com/AndreVallestero)

### Donations

If you like our distribution and have some bucks to spare, please consider contributions to the projects and developers we rely on the most:

- for sway and wlroots, consider [Drew DeVault](https://drewdevault.com/)
- for waybar, consider [Alexis Rouillard](https://github.com/sponsors/Alexays)

### Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
