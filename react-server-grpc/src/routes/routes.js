import React from "react";
import { Switch, Route, Redirect } from "react-router-dom";

export const useRoutes = () => {
	//const {token} = useContext();
	return (
		<Switch>
			<Route exact path="/" component={Home} />
		</Switch>
	);
};
