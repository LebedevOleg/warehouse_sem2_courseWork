import { createContext } from "react";
function noop() {}

export const AuthContext = createContext({
	token: null,
	userId: null,
	userEmail: null,
	userRole: null,
	isAdmin: false,
	login: noop,
	logout: noop,
});
