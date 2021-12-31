package route

const(
	REGISTER = "/register"
	LOGIN = "/login"
	ARTIKEL = "/articles"
	ARTIKEL_MANIPULATION = "/articles/{articleId}"
	COMMENT = "/articles/{articleId}/comments"
	MANIPULATION_COMMENT = "/articles/{articleId}/comments/{commentId}"
)
