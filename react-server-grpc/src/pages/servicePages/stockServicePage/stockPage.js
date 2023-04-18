import { Box, Typography } from "@mui/material";
import React from "react";
import { useAuth } from "../../../hooks/auth.hook";

const StockPage = () => {
	const { token } = useAuth();

	return (
		<>
			<Typography variant="h3">Управление Складами</Typography>
			<Box></Box>
		</>
	);
};
