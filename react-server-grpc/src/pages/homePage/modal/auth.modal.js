import {
	Button,
	Dialog,
	DialogTitle,
	Tab,
	DialogActions,
	DialogContent,
	DialogContentText,
	TextField,
} from "@mui/material";
import AccountCircle from "@mui/icons-material/AccountCircle";
import PasswordIcon from "@mui/icons-material/Password";
import { Box } from "@mui/system";
import React, { useContext } from "react";
import { AuthContext } from "../../../context/auth.context";
import TabList from "@mui/lab/TabList";
import TabContext from "@mui/lab/TabContext";
import TabPanel from "@mui/lab/TabPanel";
import axios from "axios";

const SignModal = () => {
	const [openModal, setOpenModal] = React.useState(false);
	const [tab, setTab] = React.useState("0");
	const auth = useContext(AuthContext);
	const [loginForm, setLoginForm] = React.useState({
		email: "",
		pass: "",
	});
	const [registerForm, setRegisterForm] = React.useState({
		name: "",
		email: "",
		pass: "",
	});

	const handleLogin = async () => {
		await axios
			.post(
				"http://localhost:8000/login",
				{ ...loginForm },
				{ headers: { "Content-Type": "application/json" } }
			)
			.then((res) => {
				console.log(res.data);
				auth.login(res.data.token);
				window.location.reload();
			})
			.catch((err) => {
				if (err.request) {
					console.log(err.request);
				}
			});
	};

	const handleRegister = async () => {};

	const handleOpenModal = () => {
		setOpenModal(true);
	};
	const handleCloseModal = () => {
		setOpenModal(false);
	};
	const handleChangeLogin = (e) => {
		setLoginForm({
			...loginForm,
			[e.target.name]: e.target.value,
		});
	};
	const handleChangeRegister = (e) => {
		setRegisterForm({
			...registerForm,
			[e.target.name]: e.target.value,
		});
	};
	const handleChangeTabs = (e, newValue) => {
		setTab(newValue);
	};

	return (
		<div>
			<Button color="inherit" onClick={handleOpenModal}>
				Войти/зарегистрироваться
			</Button>
			<Dialog open={openModal} onClose={handleCloseModal}>
				<DialogTitle>Вход/Регистрация</DialogTitle>

				<Box sx={{ width: "100%", typography: "body1" }}>
					<TabContext value={tab}>
						<Box sx={{ borderBottom: 1, borderColor: "divider" }}>
							<TabList
								onChange={handleChangeTabs}
								aria-label="lab API tabs example"
							>
								<Tab label="Вход" value="0" />
								<Tab label="Регистрация" value="1" />
							</TabList>
						</Box>
						<TabPanel value="0">
							<DialogContent>
								<DialogContentText>
									Введите данные для входа или перейдите на
									вкладку регистрации.
								</DialogContentText>
								<Box
									sx={{
										display: "flex",
										alignItems: "flex-end",
									}}
								>
									<AccountCircle
										sx={{
											color: "action.active",
											mr: 1,
											my: 1,
										}}
										fontSize="large"
									/>
									<TextField
										autoFocus
										name="email"
										margin="dense"
										id="email"
										label="Введите почту"
										type="email"
										fullWidth
										variant="filled"
										onChange={handleChangeLogin}
									/>
								</Box>
								<Box
									sx={{
										display: "flex",
										alignItems: "flex-end",
									}}
								>
									<PasswordIcon
										sx={{
											color: "action.active",
											mr: 1,
											my: 1,
										}}
										fontSize="large"
									/>
									<TextField
										name="pass"
										margin="dense"
										id="password"
										label="Введите пароль"
										type="password"
										fullWidth
										variant="filled"
										onChange={handleChangeLogin}
									/>
								</Box>
							</DialogContent>
							<DialogActions>
								<Button onClick={handleCloseModal}>
									Отмена
								</Button>
								<Button onClick={handleLogin}>Войти</Button>
							</DialogActions>
						</TabPanel>
						<TabPanel value="1">
							<DialogContent>
								<DialogContentText>
									Введите данные для регестрировации или
									перейдите на вкладку входа.
								</DialogContentText>
								<Box
									sx={{
										display: "flex",
										alignItems: "flex-end",
									}}
								>
									<AccountCircle
										sx={{
											color: "action.active",
											mr: 1,
											my: 1,
										}}
										fontSize="large"
									/>
									<TextField
										autoFocus
										margin="dense"
										id="firstName"
										label="Введите имя"
										type="text"
										fullWidth
										variant="filled"
										name="name"
										onChange={handleChangeRegister}
									/>
								</Box>
								<Box
									sx={{
										display: "flex",
										alignItems: "flex-end",
									}}
								>
									<AccountCircle
										sx={{
											color: "action.active",
											mr: 1,
											my: 1,
										}}
										fontSize="large"
									/>
									<TextField
										margin="dense"
										id="email"
										label="Введите почту"
										type="email"
										fullWidth
										variant="filled"
										name="email"
										onChange={handleChangeRegister}
									/>
								</Box>
								<Box
									sx={{
										display: "flex",
										alignItems: "flex-end",
									}}
								>
									<PasswordIcon
										sx={{
											color: "action.active",
											mr: 1,
											my: 1,
										}}
										fontSize="large"
									/>
									<TextField
										margin="dense"
										id="password"
										label="Введите пароль"
										type="password"
										fullWidth
										variant="filled"
										name="pass"
										onChange={handleChangeRegister}
									/>
								</Box>
							</DialogContent>
							<DialogActions>
								<Button onClick={handleCloseModal}>
									Отмена
								</Button>
								<Button onClick={handleRegister}>
									зарегистрироваться
								</Button>
							</DialogActions>
						</TabPanel>
					</TabContext>
				</Box>
			</Dialog>
		</div>
	);
};

export default SignModal;
