package repositories

type ImageRepository struct {
	// connected GORM here
}

func NewImageRepository() *ImageRepository {
	return &ImageRepository{}
}

func (repo *ImageRepository) ListAllImages() {

}
