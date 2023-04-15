import "./App.css";
import { AuthContext } from "./context/auth.context";
import { useAuth } from "./hooks/auth.hook";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import HomePage from "./pages/homePage/homePage";
import NavBar from "./components/nav-bar/nav_bar";

function App() {
	const { token, userRole, login, isAdmin, logout, ready } = useAuth();
	const isAuthent = !!token;

	if (!ready) {
		return (
			<div>
				<p>Loading...</p>
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
					</Routes>
				</BrowserRouter>
			</div>
		</AuthContext.Provider>
	);
}

export default App;
