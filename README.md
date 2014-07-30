# hosty

## Usage

### List all entries managed by hosty

    $ hosty

### Echo all /etc/hosts content

    $ hosty cat

### Create an entry

    $ sudo hosty create static 127.0.0.1 static.example.com

### Edit an entry

    $ sudo hosty edit static 127.0.0.1 static.example.com static1.example.com static2.example.com

### Enable an entry

    $ sudo hosty enable static

### Disable an entry

    $ sudo hosty disable static

### Destroy an entry

    $ sudo hosty destroy static
