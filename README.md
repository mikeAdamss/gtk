
# gtk

gtk (good to know) is a simple http request-response controller intended to allow the convenient control of large numbers (or rarely used) rest based severless functions.

It's a go wrapper around an user controlled inventory of endpoints, their descriptions and their identifers. The parts that're "good to know" is what functions are availible to you, whether a called function succeeded and what was returned in the response.

gtk assumes you'll require authentication methods and supports sha-1 http request authentication (though you can turn it off if you're feeling brave/don't care who triggers your serverless resources).


## How it works

The endpoints are defined via a user controlled inventory. Each item in the inventory consists of:

- a keyword to call the endpoint.
- a description of what it does and what it returns.
- its url.

and ..optionally
- method (if ommitted, it'll be a GET)
- source repo
- permissions*
- any headers that need to be added to the request, either hard coded or from an environmental variable.
- the chosen parser (if any), to handle the expected response.

*permissions - allows you to asign an authentication hierarchy where (for example) you have a single team function library shared between junior and senior team members.


## Installation and Setup

- Clone this repo.
- Set the env variable `GTK_INVENTORY` pointing to the included `inventory.yml`.
- (Optionally, but advised) Set the env variable `GTK_SECRET_KEY` to your chosen secret key (minimum 20 characers).
- (Optionally) set the env variable `GTK_VARS` to use a local file of variables, where using environmental variables for sending to your functions would otherwise be inconvenient/impractical.
- Navigate to the cloned repo and `go install`.

## Authentication

A hash created from your `GTK_SECRET_KEY` will be included in all outgoing requests. If no key has been provided you'll need to provide the additional flag `-na` (no authentication) from the command line.

Note - you'll still need to include server side authentication in your cloud functions. For python users there's a lightweight wrapper that provides this ...here.

## Example commands

default:
- `gtk`: lists all endpoints gtk knows about with the description of what they do and their shortform name.
- `gtk -call=<NAME>`: call the endpoint in question.

NOTE - to use the "ports" example you'll need update the url field in the example inventory.yaml.


## Defining a new inventory item

To add an endpoint to to the inventory add an entry as per the below to your `inventory.yaml` file.

```yaml
functions:

    # An example using tthree request headers, one hard coded, one taken from an environmental variables, one taken
    # from the optional json vars file.
    - name: "example"
      description: "This is just the example, a description goes here."
      url: "www.some-variation-of-serveless-function-url.com/probably-google-or-aws/example"
      headers:
         - key: "do key"
           value: "foo"
         - name: "ray key"
           key: "BAA_KEY_NAME"
         - name: "me key"
           json_key: "ME_KEY_NAME"
      response_parser: "arrayOfLines"
```

The response_parses will be build up over time as needed. The "arrayOfLines" one shown expects to be returned a response body that can be unmarashalled into an array of strings, these strings are then printed to the terminal in order.

If no `response_parser` is specified gdk will print the response status code.

If jsonParser is specified the response boy will be pretty printed to your terminal.

Response parsers are intended to be extensible and leightweight, feel free to add more to the `/parsers` package or throw a pr my way.
