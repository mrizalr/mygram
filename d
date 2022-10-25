warning: in the working copy of 'main.go', LF will be replaced by CRLF the next time Git touches it
[1mdiff --git a/handlers/comment.go b/handlers/comment.go[m
[1mindex 7aa4d73..6623eaa 100644[m
[1m--- a/handlers/comment.go[m
[1m+++ b/handlers/comment.go[m
[36m@@ -52,8 +52,6 @@[m [mfunc (h *CommentHandlers) CreateComment(c *gin.Context) {[m
 }[m
 [m
 func (h *CommentHandlers) GetAllComment(c *gin.Context) {[m
[31m-	c.MustGet("claims")[m
[31m-[m
 	comments, err := h.commentService.GetAll()[m
 	if err != nil {[m
 		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{[m
[1mdiff --git a/handlers/socialMedia.go b/handlers/socialMedia.go[m
[1mindex e2c51cd..5313654 100644[m
[1m--- a/handlers/socialMedia.go[m
[1m+++ b/handlers/socialMedia.go[m
[36m@@ -24,7 +24,7 @@[m [mfunc (h *SocialMediaHandlers) AddSocialMedia(c *gin.Context) {[m
 	if err := c.ShouldBindJSON(&socmedRequest); err != nil {[m
 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{[m
 			"status": "error",[m
[31m-			"error":  err.Error(),[m
[32m+[m			[32m"message":  err.Error(),[m
 		})[m
 		return[m
 	}[m
[36m@@ -36,10 +36,23 @@[m [mfunc (h *SocialMediaHandlers) AddSocialMedia(c *gin.Context) {[m
 	if err != nil {[m
 		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{[m
 			"status": "error",[m
[31m-			"error":  err.Error(),[m
[32m+[m			[32m"message":  err.Error(),[m
 		})[m
 		return[m
 	}[m
 [m
 	c.JSON(http.StatusCreated, socialMedia)[m
 }[m
[32m+[m
[32m+[m[32mfunc (h *SocialMediaHandlers) GetAllSocmeds(c *gin.Context) {[m
[32m+[m	[32msocmeds, err := h.socmedService.GetAll()[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mc.AbortWithStatusJSON(http.StatusBadGateway, gin.H{[m
[32m+[m			[32m"status": "error",[m
[32m+[m			[32m"message":  err.Error(),[m
[32m+[m		[32m})[m
[32m+[m		[32mreturn[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mc.JSON(http.StatusOK, socmeds)[m
[32m+[m[32m}[m
[1mdiff --git a/main.go b/main.go[m
[1mindex 6843f9d..2044676 100644[m
[1m--- a/main.go[m
[1m+++ b/main.go[m
[36m@@ -35,7 +35,7 @@[m [mfunc main() {[m
 	routers.InitCommentRoutes(Routes, commentHandler)[m
 [m
 	socmedRepository := repositories.NewSocmedRepository(database.GetDB())[m
[31m-	socmedService := services.NewSocmedService(socmedRepository)[m
[32m+[m	[32msocmedService := services.NewSocmedService(socmedRepository, userRepository)[m
 	socmedHandler := handlers.NewSocialMediaHandlers(socmedService)[m
 	routers.InitSocmedRouter(Routes, socmedHandler)[m
 [m
[1mdiff --git a/models/socialMedia.go b/models/socialMedia.go[m
[1mindex 8eded2a..3790eaf 100644[m
[1m--- a/models/socialMedia.go[m
[1m+++ b/models/socialMedia.go[m
[36m@@ -1,6 +1,46 @@[m
 package models[m
 [m
[32m+[m[32mimport ([m
[32m+[m	[32m"time"[m
[32m+[m
[32m+[m	[32m"github.com/mrizalr/mygram/entities"[m
[32m+[m[32m)[m
[32m+[m
 type AddSocialMediaRequest struct {[m
 	Name           string `json:"name" binding:"required"`[m
 	SocialMediaURL string `json:"social_media_url" binding:"required"`[m
 }[m
[32m+[m
[32m+[m[32mtype GetSocmedResponse struct {[m
[32m+[m	[32mID             int       `json:"id"`[m
[32m+[m	[32mName           string    `json:"name"`[m
[32m+[m	[32mSocialMediaURL string    `json:"social_media_url"`[m
[32m+[m	[32mUserID         int       `json:"user_id"`[m
[32m+[m	[32mCreatedAt      time.Time `json:"created_at"`[m
[32m+[m	[32mUpdatedAt      time.Time `json:"updated_at"`[m
[32m+[m	[32mUser           struct {[m
[32m+[m		[32mID       int    `json:"id"`[m
[32m+[m		[32mUsername string `json:"username"`[m
[32m+[m		[32mEmail    string `json:"email"`[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc ParseToGetSocmedResponse(socmed entities.SocialMedia, user entities.User) GetSocmedResponse {[m
[32m+[m	[32mreturn GetSocmedResponse{[m
[32m+[m		[32mID:             int(socmed.ID),[m
[32m+[m		[32mName:           socmed.Name,[m
[32m+[m		[32mSocialMediaURL: socmed.SocialMediaURL,[m
[32m+[m		[32mUserID:         int(socmed.UserID),[m
[32m+[m		[32mCreatedAt:      socmed.CreatedAt,[m
[32m+[m		[32mUpdatedAt:      socmed.UpdatedAt,[m
[32m+[m		[32mUser: struct {[m
[32m+[m			[32mID       int    `json:"id"`[m
[32m+[m			[32mUsername string `json:"username"`[m
[32m+[m			[32mEmail    string `json:"email"`[m
[32m+[m		[32m}{[m
[32m+[m			[32mID:       int(user.ID),[m
[32m+[m			[32mUsername: user.Username,[m
[32m+[m			[32mEmail:    user.Email,[m
[32m+[m		[32m},[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[1mdiff --git a/repositories/socialMedia.go b/repositories/socialMedia.go[m
[1mindex 42faeb8..1eee311 100644[m
[1m--- a/repositories/socialMedia.go[m
[1m+++ b/repositories/socialMedia.go[m
[36m@@ -1,12 +1,13 @@[m
 package repositories[m
 [m
 import ([m
[31m-	"gorm.io/gorm"[m
 	"github.com/mrizalr/mygram/entities"[m
[32m+[m	[32m"gorm.io/gorm"[m
 )[m
 [m
 type SocmedRepository interface {[m
 	Create(socmed entities.SocialMedia) (entities.SocialMedia, error)[m
[32m+[m	[32mFindAll() ([]entities.SocialMedia, error)[m
 }[m
 [m
 type socmedRepository struct {[m
[36m@@ -21,4 +22,11 @@[m [mfunc NewSocmedRepository(db *gorm.DB) *socmedRepository {[m
 [m
 func (r *socmedRepository) Create(socmed entities.SocialMedia) (entities.SocialMedia, error) {[m
 	return socmed, r.db.Create(&socmed).Error[m
[31m-}[m
\ No newline at end of file[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (r *socmedRepository) FindAll() ([]entities.SocialMedia, error){[m
[32m+[m	[32msocialMedias := []entities.SocialMedia{}[m
[32m+[m	[32merr := r.db.Find(&socialMedias).Error[m
[32m+[m
[32m+[m	[32mreturn socialMedias, err[m
[32m+[m[32m}[m
[1mdiff --git a/routers/socialMedia.go b/routers/socialMedia.go[m
[1mindex 4325808..c756191 100644[m
[1m--- a/routers/socialMedia.go[m
[1m+++ b/routers/socialMedia.go[m
[36m@@ -10,5 +10,6 @@[m [mfunc InitSocmedRouter(Routes *gin.Engine, handler *handlers.SocialMediaHandlers)[m
 	socmedGroup := Routes.Group("/socialmedias")[m
 	{[m
 		socmedGroup.POST("/", middlewares.Auth, handler.AddSocialMedia)[m
[32m+[m		[32msocmedGroup.GET("/", middlewares.Auth, handler.GetAllSocmeds)[m
 	}[m
 }[m
[1mdiff --git a/services/socialMedia.go b/services/socialMedia.go[m
[1mindex 23e9345..f4fe7c9 100644[m
[1m--- a/services/socialMedia.go[m
[1m+++ b/services/socialMedia.go[m
[36m@@ -8,15 +8,18 @@[m [mimport ([m
 [m
 type SocmedService interface {[m
 	Add(userID int, socmed models.AddSocialMediaRequest) (entities.SocialMedia, error)[m
[32m+[m	[32mGetAll() ([]models.GetSocmedResponse, error)[m
 }[m
 [m
 type socmedService struct {[m
 	socmedRepository repositories.SocmedRepository[m
[32m+[m	[32muserRepository   repositories.UserRepository[m
 }[m
 [m
[31m-func NewSocmedService(socmedRepository repositories.SocmedRepository) *socmedService {[m
[32m+[m[32mfunc NewSocmedService(socmedRepository repositories.SocmedRepository, userRepository repositories.UserRepository) *socmedService {[m
 	return &socmedService{[m
 		socmedRepository: socmedRepository,[m
[32m+[m		[32muserRepository:   userRepository,[m
 	}[m
 }[m
 [m
[36m@@ -29,3 +32,22 @@[m [mfunc (s *socmedService) Add(userID int, socmedRequest models.AddSocialMediaReque[m
 [m
 	return s.socmedRepository.Create(socmed)[m
 }[m
[32m+[m
[32m+[m[32mfunc (s *socmedService) GetAll() ([]models.GetSocmedResponse, error) {[m
[32m+[m	[32mresponse := []models.GetSocmedResponse{}[m
[32m+[m	[32msocmeds, err := s.socmedRepository.FindAll()[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn response, err[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mfor _, socmed := range socmeds {[m
[32m+[m		[32muser, err := s.userRepository.FindByID(int(socmed.UserID))[m
[32m+[m		[32mif err != nil {[m
[32m+[m			[32mreturn response, err[m
[32m+[m		[32m}[m
[32m+[m
[32m+[m		[32mresponse = append(response, models.ParseToGetSocmedResponse(socmed, user))[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mreturn response, nil[m
[32m+[m[32m}[m
