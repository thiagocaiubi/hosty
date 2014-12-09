# DANGER: Lots of things missing! Use it at your own risk, or help me get things done! ;)

# hosty

## Usage

### List all entries managed by hosty

    $ hosty

### Echo all /etc/hosts content

    $ hosty cat

### Save an entry, use this to create or edit an entry

    $ sudo hosty save static 127.0.0.1 static.example.com

### Enable an entry

    $ sudo hosty enable static

### Disable an entry

    $ sudo hosty disable static

### Remove an entry

    $ sudo hosty remove static
