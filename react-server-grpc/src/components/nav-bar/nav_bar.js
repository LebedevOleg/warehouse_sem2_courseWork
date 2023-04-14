import React from "react";
import Box from "@mui/material/Box";
import {
	AppBar,
	Button,
	IconButton,
	Menu,
	MenuItem,
	Toolbar,
	Typography,
} from "@mui/material";
import { useAuth } from "../../hooks/auth.hook";
import SignModal from "../../pages/homePage/modal/auth.modal";

const NavBar = () => {
	const { token } = useAuth();

	return (
		<>
			<Box sx={{ flexGrow: 1 }}>
				<AppBar position="static">
					<Toolbar>
						<Typography
							variant="h6"
							component="div"
							sx={{ flexGrow: 1 }}
						>
							Склад
						</Typography>
						<Typography>
							<Button id="basic-button">Аккаунт</Button>
							<Menu>
								<MenuItem id="account">Аккаунт</MenuItem>
								<MenuItem id="service">Сервисное меню</MenuItem>
							</Menu>
						</Typography>
						{(!!token && (
							<Typography>ВЫ ЗАЛОГИНЕНЫ</Typography>
						)) || <SignModal />}
					</Toolbar>
				</AppBar>
			</Box>
		</>
	);
};

export default NavBar;
