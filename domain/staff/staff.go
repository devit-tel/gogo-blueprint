package staff

import (
    "github.com/devit-tel/gotime"
)

type Staff struct {
    Id        string `bson:"_id"`
    CompanyId string `bson:"companyId"`
    Name      string `bson:"name"`
    Tel       string `bson:"tel"`
    CreatedAt int64  `bson:"createdAt"`
    UpdatedAt int64  `bson:"updatedAt"`
}

func Create(id, companyId, name, tel string) *Staff {
    return &Staff{
        Id:        id,
        CompanyId: companyId,
        Name:      name,
        Tel:       tel,
        CreatedAt: gotime.NowUnix(),
        UpdatedAt: gotime.NowUnix(),
    }
}

func (s *Staff) Update(name, tel string) {
    s.Name = name
    s.Tel = tel
    s.UpdatedAt = gotime.NowUnix()
}
