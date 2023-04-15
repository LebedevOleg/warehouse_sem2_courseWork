import { useCallback, useEffect, useState } from "react";

const storage = "userData";

export const useAuth = () => {
	const [token, setToken] = useState(null);
	const [userId, setUserId] = useState(null);
	const [userEmail, setUserEmail] = useState(null);
	const [userRole, setUserRole] = useState(null);
	const [isAdmin, setIsAdmin] = useState(false);
	const [ready, setReady] = useState(false);

	const login = useCallback((t, role) => {
		setUserRole(role);
		setIsAdmin(role === "Администратор");
		setToken(t);
		localStorage.setItem(storage, JSON.stringify({ token: t }));
	}, []);

	const logout = useCallback(() => {
		setToken(null);
		setUserRole(null);
		setIsAdmin(false);
		localStorage.removeItem(storage);
	}, []);

	useEffect(() => {
		const data = JSON.parse(localStorage.getItem(storage));
		if (data && data.token) {
			login(data.token, data.role);
		}
		setReady(true);
	}, [login]);

	return {
		token,
		userRole,
		isAdmin,
		ready,
		login,
		logout,
	};
};
