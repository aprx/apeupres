## Apeupres

Apeupres is a tool for managing environment variables in a simple, shareable, and reproducible way.
It lets you define your environment variables as code, making it easy to switch between configurations (such as development, staging, or production) and share them with your team.

`Environ`, well, "roughly" means `À peu près` in French, hence the name.

----

## Usage

```sh
$ apeupres -h
Usage of ./apeupres:
  -config string
    	Path to configuration (default "/home/$HOME/.config/apeupres")
  -output string
    	Output path (default "/home/$HOME/.apeupresrc")
```

## Example

If you have a command-line application that requires `LOGIN` and `PASSWORD` environment variables to run.

You can define these variables in `~/.config/apeupres/configuration.ini` like this:

```
[production]
LOGIN=prod-login
PASSWORD=prod-password
```

byB running  `apeupres`, this will generates a shell function for you in `~/.apeupresrc`

```
conf_production() {
  apeupres_set_env production
  apeupres_set_clean_env LOGIN:PASSWORD
  export LOGIN=admin
  export PASSWORD=coucou4242
}
```

You can then run `source ~/.apeupresrc` and run `conf_production`. It sets up your shell with the correct environment variables for the production environment. `~/.apeupresrc` is overwritten at every call and should not be modified by hand.

```sh
$ conf_production
$  echo "->$LOGIN $PASSWORD<-"
->admin coucou4242<-
$ conf_unset
$ echo $LOGIN $PASSWORD
-> <-

```

### Using a shared secret vault

If you store secrets in a shared vault (for example, 1Password), you can reference them directly:
```
PASSWORD=`op item get prod-password --format json | jq .id`
```
or gopass
```
PASSWORD=`gopas show /path/to/secret`
```

This way, sensitive values never need to be stored in plain text.


## Configuration

By default, Apeupres:

- Looks for *every* .ini files in
```
$HOME/.config/apeupres/
```

- Generates its output in
```
$HOME/.apeupresrc
```

To enable Apeupres in your shell, add the following line to your shell configuration file (e.g. .bashrc, .zshrc):
```
source ~/.apeupresrc
```
### Ini files

`Apeupres` will use the section name to generate the shell function by prefixing it with `conf_<SECTION>`. Defining 2 sections with the same name will result in unknwon behavior (well last defined will override the first previous one).

Also it's not possible to name a section `DEFAULT` or `unset`, the section will be ignored or generate an error.



That’s it—you’re ready to use Apeupres.

## Shell prompt integration

Apeupres sets the `APEUPRES_NAME` environment variable to indicate which configuration is currently active.
You can use this to display the active environment directly in your shell prompt.

```sh
PS1='${APEUPRES_NAME:+[|$APEUPRES_NAME|]}\u@\h:\w\$ '
```

Here’s how to show the active Apeupres environment in a [Starship](https://starship.rs/config/#environment-variable) prompt:
```toml
[env_var.APEUPRES_NAME]
variable = "APEUPRES_NAME"
format = '\[|[$env_value](red)|\] '
```



