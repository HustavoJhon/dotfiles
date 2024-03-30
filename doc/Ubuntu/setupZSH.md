# This is what i every time do After installing a ubuntu/debian on Physical/Virtual computer to setup ZSH.

**Check Shell**

```bash
echo %0
```
1. Install Git, Curl, Wget as prequisites

```bash
sudo apt install curl git wget
```

2. Install ZSH shell

```bash
sudo apt install zsh
```

3. Change Default Shell

```bash
chsh -s $(which zsh)
```

#### -----> Now Reboot The Computer (linux) <-----

**Check Shell Again**

```bash
echo $0
```

4. Download [Oh-My-ZSH](https://ohmyz.sh/)

```bash
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```
5. Download any [Powerline Font](https://github.com/powerline/fonts) -> [Fira font](https://github.com/tonsky/FiraCode), [Hack Font](https://github.com/source-foundry/Hack) or other to get best experience.

6. Download [ZSH Auto Suggestion](https://github.com/zsh-users/zsh-autosuggestions)

```bash
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
```

7. Download [ZSH Syntax Highlighting](https://github.com/zsh-users/zsh-syntax-highlighting)

```bash
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
```

8. Add this to ~/.zshrc file

**Obviously before `source $ZSH/oh-my-zsh.sh` part**

```bash
# default theme
# ZSH_THEME="robbyrussell"
ZSH_THEME="agnoster"

plugins=(
    git
    history
    zsh-autosuggestions
    zsh-syntax-highlighting
    command-not-found
)
```
9. (Optional) Download Power Level 10k Theme

```bash
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k
echo 'source ~/powerlevel10k/powerlevel10k.zsh-theme' >>~/.zshrc
```
10. (If you chose 9) Go to powerlevel10k folder & run this to customize style

```bash
p10k configure
```
