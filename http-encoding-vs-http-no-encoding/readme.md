## https vs axios 


| site  | encoding  | no encoding  | response size (encoding) | response size  |
|---|---|---|---|---|
| example.com | `533.5298489992321` |  `536.7082909996808` |  850B |  1.3KB |
| html.spec | `2795.3802000015976` |  `7161.993405000865` |  1.9MB |  12.9MB |

### https with encoding ( [example.com](https://example.com) ) 1.3KB --> 850B
https after running it 100 times `533.5298489992321` ms

### https without encoding ( [example.com](https://example.com) ) 1.3KB
after running it 100 times `536.7082909996808` ms

### https with encoding ( [html.spec](https://html.spec.whatwg.org/) ) 12.9MB --> 1.9MB
after running it 20 times `2795.3802000015976` ms

### https without encoding (  [html.spec](https://html.spec.whatwg.org/) ) 12.9MB
https after running it 20 times `7161.993405000865` ms

## thoughts
encoding is a magic

# See also 
[accept-encoding header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding)