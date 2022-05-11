package service

// OrderService 处理订单的服务
type OrderService struct {
	CusSvc *CustomerService
}

func NewOrderService(cusSvc *CustomerService) *OrderService {
	return &OrderService{CusSvc: cusSvc}
}

func (svc *OrderService) CreateOrder(cid int, gid int, num int) {
	goods, err := svc.GetGoodsById(gid)
	if err != nil {
		// 处理错误
		return
	}
	total := num * goods.Price
	// 扣钱
	idx := svc.gsvc.FindById(cid)
	if idx < 0 {
		// 打印错误
		return
	}
	svc.gsvc[idx].Money -= total

	// 创建订单

}
