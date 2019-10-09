package nameservice

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case MsgAddNewCoin:
			return handleMsgAddNewCoin(ctx, keeper, msg)
		case MsgAddCoin:
			return handleMsgAddCoin(ctx, keeper, msg)
		case MsgBurnCoin:
			return handleMsgBurnCoin(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set name
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                      // return
}

// Handle a message to buy name
func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) { // Checks if the the bid price is greater than the price paid by the current owner
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}
	if keeper.HasOwner(ctx, msg.Name) {
		err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		_, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return sdk.Result{}
}


// ---------------------------------------------
// Add new token to account
func handleMsgAddNewCoin(ctx sdk.Context, keeper Keeper, msg MsgAddNewCoin) sdk.Result {
	// 限定一次只发行一个，不能和之前的重复
	oldCoins := keeper.coinKeeper.GetCoins(ctx, msg.Owner)
	newCoins := msg.Amt
	var addCoins sdk.Coins

	addCoin := newCoins[0]
	addCoins = append(addCoins, addCoin)

	// 判断重复, 如果要新发行的coin和原来的重复，该怎么输出错误信息？
	in := contain(addCoin, oldCoins)
	if !in {
		_, err := keeper.AddNewCoins(ctx, msg.Owner, addCoins)
		if err != nil {
			return sdk.ErrUnknownRequest("Owner address invalid or coins invalid").Result()
		}
	}

	return sdk.Result{}
}

// 增发
func handleMsgAddCoin(ctx sdk.Context, keeper Keeper, msg MsgAddCoin) sdk.Result {
	//oldCoins := keeper.coinKeeper.GetCoins(ctx, msg.Owner)
	newCoins := msg.Amt
	var addCoins sdk.Coins

	addCoin := newCoins[0]
	addCoins = append(addCoins, addCoin)

	// 增发，不用判断重复
	_, err := keeper.AddCoin(ctx, msg.Owner, addCoins)
	if err != nil {
		return sdk.ErrUnknownRequest("Owner address invalid or coins invalid").Result()
	}

	return sdk.Result{}
}

// 销毁token,

// Burn a token from account
func handleMsgBurnCoin(ctx sdk.Context, keeper Keeper, msg MsgBurnCoin) sdk.Result {
	oldCoins := keeper.coinKeeper.GetCoins(ctx, msg.Owner)
	amount := msg.Amt
	burnCoin := amount[0]
	var burnCoins sdk.Coins
	burnCoins = append(burnCoins, burnCoin)

	//一次销毁一个币,从切片中删除某个Coin,销毁之前要先判断该账户下有没有这个币,还要判断要销毁的币的数量是否超出原来账户下该币的数量
	// 如果要销毁的币的数量 > 原来账户下该币的数量，该怎么输出错误信息？
	// ok1,********
	//ok is:false

	ok := safeBurn(burnCoin, oldCoins)
	fmt.Println("ok1,********")
	fmt.Printf("ok is:%v \n", ok)
	if ok {
		fmt.Println("ok2,********")
		_, err := keeper.BurnCoins(ctx, msg.Owner, burnCoins)
		if err != nil {
			return sdk.ErrUnknownRequest("Owner address invalid or coins invalid").Result()
		}
	}

	return sdk.Result{}
}

func contain(Coin sdk.Coin, Coins sdk.Coins) bool {
	for i := 0; i < Coins.Len(); i++ {
		if Coins[i].Denom == Coin.Denom {
			return true
		}
	}
	return false
}

//func safeBurn(Coin sdk.Coin, Coins sdk.Coins) bool {
//	for i := 0; i < Coins.Len(); i++ {
//		if Coins[i].Denom == Coin.Denom && Coins[i].Amount == Coin.Amount {
//			return true
//		}
//	}
//	return false
//}

func safeBurn(Coin sdk.Coin, Coins sdk.Coins) bool {
	for i := 0; i < Coins.Len(); i++ {
		if Coins[i].Denom == Coin.Denom {
			return true
		}
	}
	return false
}









