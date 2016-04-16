Minimal PERC Demo
=================

```
          KMF
 _____________________
/                     \

+------+     +--------+     +-----+                  +----------+
| DTLS |-udp-| Tunnel |=====| MDD |---DTLS/SRTP--//--| Browser+ |
+------+     +--------+     +-----+                  +----------+
    .            ^
    .    keys    .
    ..............

```

* Implement a trivial MDD in Go
  * Configuration:
    * Tunnel port
    * UDP port(s) for media
  * On UDP packet
    * If DTLS, send down tunnel, note peer it came from
    * If UDP, re-encrypt, route to others
  * On tunnel packet
    * Read crypto info if present
    * Forward to appropriate UDP peer

* KMF is a hybrid of Go and an existing DTLS stack
  * Implement -tunnel endpoint in Go, complement of MDD side
  * Pass DTLS packets over UDP ports
    * Use a new UDP port for each -tunnel connection ID
    * ... to convince the DTLS stack it's a new connection
  * Need a way for DTLS to exfiltrate keys to -tunnel
    * NSS and OpenSSL support export to SSLKEYLOGFILE
    * Maybe that's enough?
    * Could probably get DTLS handshake done that way, not EKT

* Browser is just an existing browser, plus -double support
  * Feed it some crafted SDP to get it to connect to the KMF via MDD
  * Bonus points for IdP

* DTLS stack and Browser probably need to support EKT
  * Could do an initial demo with one endpoint and no EKT
  * But that's not so interesting
