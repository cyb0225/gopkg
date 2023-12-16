# go pkg

Some simple Go toolkits to help speed up daily development.

Some of them were copied from well-known projects because I think their dependencies are too extensive, and many of them are not useful to me.
Or I have made modifications to these codes, such as using generics

- **bizerror** encapsulates code and msg into errors, providing asynchronous error callbacks and error stack printing.
- **context**
  - **without_cancel** cancels the parent context but keeps keys and values, consistent with WithoutCancel in go1.21
- **copy**
  - **deepcopy** supports deep copy of generics, implemented based on reflection
- **grpc**
  - **middleware-chain**
  - **context-wrappers**
- **logs** which You can use it directly without initialization
  - **zap logger**
- **routine** opens a goroutine to recover the panic automatically


