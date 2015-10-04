# Networking with Go lang

##Why Go?
- Go has some advantages over Java and Python
- Simpler, faster and lighter
- Advanced networking and database libraries available, such as etcd
- I want to rewrite part of my exisiting programs ([neutron-lan](https://github.com/araobp/neutron-lan), [tega](https://github.com/araobp/tega)) in Go lang with its libraries

```
[Networking agent written in Go]
          |
    Cross compile
          |
    +-----+-----+
    |           |
    V           V
[CPU:ARM]   [CPU:x86]
```

##Preparations
- Go lang installation: https://golang.org/dl/
- Protobuf build and installation: https://github.com/google/protobuf/blob/master/INSTALL.txt
```
$ autoconf
$ ./autogen.sh
$ ./configure
$ make
$ make install
```
- Add /usr/local/lib to LD_LIBRARY_PATH
```
$ export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARLY_PATH
```
##Go plugin for vim

Install [vim-go](https://github.com/fatih/vim-go) to your vim.

##Interesting libraries written in Go lang
- [gRPC](https://github.com/grpc/grpc-go/)
- [etcd](https://github.com/coreos/etcd)
- [OVSDB client](https://github.com/socketplane/libovsdb)
- [YAML](https://github.com/go-yaml/yaml)

##Reference
- [Software Defined 
Networking at Scale, Bikash Koley, Anees Shaikh on behalf of Google Technical Infrastructure
, 5/12/2015](http://files.meetup.com/8218762/Bikash_Koley%20SDN_meetup%20May%202015.pdf)