package global

const (
	dburl       = "mongodb+srv://standard:passw0rd@mycluster-l25ub.mongodb.net/test?retryWrites=true&w=majority"
	dbname      = "blog-application"
	performance = 100
)

var (
	jwtSecret = []byte("blogSecret")
)
