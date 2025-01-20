# geopass - CLI Password Generator

<img src="images/geopass-logo.png" alt="geopass Logo" width="200"/>

## About

`geopass` is a simple command-line tool for generating strong passwords. Written in Go, it provides users with the ability to quickly generate random, secure passwords directly from the terminal. Whether you need a password for a new account, API key, or anything requiring security, geopass can generate complex passwords with customizable options.

## Commands

`geopass` CLI offers the following commands:

| Command | Description                                                                                                                                              | Example Command           |
| ------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------- |
| `-v`    | Display the current version of geopass.                                                                                                                  | `geopass -v`              |
| `gen`   | Generate a new secure password. The password is randomly created with a default length or specified length (optional). `-l <length> or -length <length>` | `geopass gen --length=25` |

## How To Use

**_NOTE:_** - It is recommended to use `Docker` to run geopass CLI.

1. Clone the repo

```bash
git clone https://github.com/georgelopez7/geopass-cli.git
```

2. Build Docker image

```bash
docker build -t geopass-cli .
```

**_NOTE_**: You will notice that the image contains `tail -f /dev/null` in the Dockerfile. This is to prevent the container from exiting immediately starting the command.

3. Run Docker container

```bash
docker run -d --name geopass-container geopass-cli
```

4. Execute geopass CLI commands

```bash
# Check geopass version
docker exec geopass-container geopass-cli geopass --version

# Generate a new password
docker exec geopass-container geopass-cli geopass gen --length=25
```

## Example output:

<pre>
<code>
<span style="color:white;">Generated Password:</span>
<span style="color:cyan;">z).}diStrSPfb`b-GnF,`;8~# </span>

<span style="color:white;">Password Entropy:</span>
<span style="color:lightgreen;">162.69</span>
</code>
</pre>

## What is Entropy?

Entropy measures the randomness or unpredictability of a password. In the context of password generation, higher entropy means a stronger, more secure password, while lower entropy indicates weaker security.

Shannon entropy is used to calculate password strength based on its length and the number of possible characters (pool size). In the code, this is done by multiplying the password length by the information provided by each character, `log2(pool size)`.

Entropy helps ensure that generated passwords are robust and resistant to guessing or cracking.
