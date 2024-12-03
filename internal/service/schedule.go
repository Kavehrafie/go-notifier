package service

//type ScheduleCallback func(*model.ScheduledAction)
//type ScheduleNotificationCallback func(*model.ScheduledAction, int)
//
//type ScheduleService interface {
//	// schedule management
//	RegisterSchedule(ctx context.Context, input *model.ScheduleActionRegisterInput) (*model.ScheduledAction, error)
//	GetSchedule(ctx context.Context, id string) (*model.ScheduledAction, error)
//	DeleteSchedule(ctx context.Context, id string) error
//
//	// Schedule operations
//	PauseSchedule(ctx context.Context, id string) (*model.ScheduledActionStatus, error)
//	ResumeSchedule(ctx context.Context, id string) (*model.ScheduledActionStatus, error)
//	TriggerSchedule(ctx context.Context, id string) (*model.ScheduledActionStatus, error)
//
//	// Callback registration
//	OnScheduleRegistered(callback ...ScheduleCallback)
//	OnScheduleResumed(callback ...ScheduleCallback)
//	OnSchedulePaused(callback ...ScheduleCallback)
//	OnScheduleNotified(callback ...ScheduleNotificationCallback)
//}
//
//type scheduleService struct {
//	store           store.Store
//	notificationSvc NotificationService
//
//	// Callbacks for different events
//	onScheduleRegistered []ScheduleCallback
//	onSchedulePaused     []ScheduleCallback
//	onScheduleResumed    []ScheduleCallback
//	onScheduleNotified   []ScheduleNotificationCallback
//}
//
//func NewScheduleService(store store.Store, notificationSvc NotificationService) ScheduleService {
//	return &scheduleService{
//		store:           store,
//		notificationSvc: notificationSvc,
//	}
//}
