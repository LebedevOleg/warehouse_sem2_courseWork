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
import { Box } from "@mui/material";

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
							path="/service"
							element={<ServicePage />}
						/>
					</Routes>
				</BrowserRouter>
			</div>
		</AuthContext.Provider>
	);
}

export default App;
