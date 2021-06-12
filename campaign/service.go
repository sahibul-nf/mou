package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
	FindCampaign(input GetCampaignDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {

	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)

		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) FindCampaign(input GetCampaignDetailInput) (Campaign, error) {

	ID := input.ID

	campaign, err := s.repository.FindByID(ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
