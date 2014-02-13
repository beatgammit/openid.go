This is a consumer (Relying party) implementation of OpenId 2.0,
written in go.

Example
========
See openid/openid_example for a simple webserver using the openID
implementation with FAS as service provider. Also, read the comment about the NonceStore towards
the top of that file.

`openid.Verify` now returns a map[string]string with keys
`user,nick,timezone,fullname,email,teams`.

License
=======

Distributed under the Apache v2.0 license:
http://www.apache.org/licenses/LICENSE-2.0.html

