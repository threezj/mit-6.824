package raftkv

const (
	OK        = "OK"
	ErrNoKey  = "ErrNoKey"
	ERROR     = "ERROR"
	ErrDupReq = "ErrDupReq"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	Key   string
	Value string
	Op    string
	ReqId int64
	Id    int64

	// "Put" or "Append"
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
}

type PutAppendReply struct {
	WrongLeader bool
	Err         Err
}

type GetArgs struct {
	Key string
	// You'll have to add definitions here.
	ReqId int64
	Id    int64
}

type GetReply struct {
	WrongLeader bool
	Err         Err
	Value       string
}
