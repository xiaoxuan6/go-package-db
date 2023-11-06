package go_package_db

type Collect struct {
	ID   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"varchar(125);not null;comment:'第三方包名'"`
	Url  string `json:"url" gorm:"varchar(125);not null;comment:'第三方包地址'"`
}

func Insert(collect ...Collect) error {
	if err := DB.Model(&Collect{}).Create(&collect).Error; err != nil {
		return err
	}
	return nil
}

func FindByName(name string) (collect Collect, err error) {
	err = DB.Model(&Collect{}).Where("name = ?", name).First(&collect).Error
	return
}
