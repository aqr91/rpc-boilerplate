package auth

type LoginArgs struct {
	Username string
	Password string
}

type LoginReply struct {
	Message string
	Token   string
}

type RegisterArgs struct {
	Username string
	Password string
}

type RegisterReply struct {
	Message string
}
