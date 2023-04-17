import React, { useContext, useState } from "react";
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
import SignModal from "../authModal/auth.modal";
import { AuthContext } from "../../context/auth.context";

const NavBar = () => {
	const { token, isAdmin } = useAuth();
	const auth = useContext(AuthContext);
	const [openProf, setOpenProf] = useState(null);
	const open = Boolean(openProf);

	const handleClick = (event) => {
		setOpenProf(event.currentTarget);
	};
	const handleClose = (event) => {
		switch (event.target.id) {
			case "logout":
				auth.logout();
				setOpenProf(null);
				window.location = "/";
				break;
			case "service":
				window.location = "/service";
				break;
			case "accaunt":
				window.location = "/profile";
				break;
			default:
				setOpenProf(null);
				break;
		}
	};

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
							<Button
								sx={{ my: 2, color: "white", display: "block" }}
								onClick={() => {
									window.location = "/shop";
								}}
							>
								Все товары
							</Button>
						</Typography>
						<Typography>
							<Button id="basic-button">Аккаунт</Button>
							<Menu>
								<MenuItem id="account">Аккаунт</MenuItem>
								<MenuItem id="service">Сервисное меню</MenuItem>
							</Menu>
						</Typography>
						{(!!token && (
							<Typography>
								<Button
									id="basic-button"
									aria-controls={
										open ? "basic-menu" : undefined
									}
									aria-haspopup="true"
									aria-expanded={open ? "true" : undefined}
									onClick={handleClick}
									sx={{
										my: 2,
										color: "white",
										display: "block",
									}}
								>
									Профиль
								</Button>
								<Menu
									id="basic-menu"
									anchorEl={openProf}
									open={open}
									onClose={handleClose}
									MenuListProps={{
										"aria-labelledby": "basic-button",
									}}
								>
									<MenuItem
										id="accaunt"
										onClick={handleClose}
									>
										Мой аккаунт
									</MenuItem>
									{isAdmin && (
										<MenuItem
											id="service"
											onClick={handleClose}
										>
											Сервисное меню
										</MenuItem>
									)}

									<MenuItem id="logout" onClick={handleClose}>
										Выйти
									</MenuItem>
								</Menu>
							</Typography>
						)) || <SignModal />}
					</Toolbar>
				</AppBar>
			</Box>
		</>
	);
};

export default NavBar;
