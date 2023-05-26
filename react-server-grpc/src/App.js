import "./App.css";
import { AuthContext } from "./context/auth.context";
import { useAuth } from "./hooks/auth.hook";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import HomePage from "./pages/homePage/homePage";
import NavBar from "./components/nav-bar/nav_bar";
import ShopPage from "./pages/shopPage/shopPage";
import ServicePage from "./pages/servicePages/servicePage";
import ItemServicePage from "./pages/servicePages/itemsServicePage/itemServicePage";
import Loader from "./components/Loading/Loading";
import DeliveryPage from "./pages/servicePages/deliveryPage/deliveryPage";
import UserServicePage from "./pages/servicePages/usersServicePage/userServicePage";
import BacketPage from "./pages/backetPage/backetPage";
import OrdersPage from "./pages/servicePages/ordersPage/ordersPage";
import StockPage from "./pages/servicePages/stockServicePage/stockPage";
import UserPage from "./pages/userPage/userPage";

function App() {
	const { token, userRole, login, isAdmin, logout, ready } = useAuth();
	const isAuthent = !!token;

	if (!ready) {
		return (
			<div>
				<Loader />
			</div>
		);
	}

	return (
		<AuthContext.Provider
			value={{
				isAdmin,
				token,
				userRole,
				login,
				logout,
				isAuthent,
			}}
		>
			<NavBar />
			<div className="container">
				<BrowserRouter>
					<Routes>
						<Route exact path="/" element={<HomePage />} />
						<Route exact path="/shop" element={<ShopPage />} />
						<Route
							exact
							path="/service/items"
							element={<ItemServicePage />}
						/>
						<Route
							exact
							path="/service/stocks"
							element={<StockPage />}
						/>
						<Route
							exact
							path="/service/delivery"
							element={<DeliveryPage />}
						/>
						<Route
							exact
							path="/service/users"
							element={<UserServicePage />}
						/>
						<Route
							exact
							path="/service"
							element={<ServicePage />}
						/>
						<Route exact path="/backet" element={<BacketPage />} />
						<Route
							exact
							path="/service/orders"
							element={<OrdersPage />}
						/>
						<Route
							exact
							path="/service/providers"
							element={<ServicePage />}
						/>
						<Route exact path="/profile" element={<UserPage />} />
						<Route path="*" element={<HomePage />} />
					</Routes>
				</BrowserRouter>
			</div>
		</AuthContext.Provider>
	);
}

export default App;
