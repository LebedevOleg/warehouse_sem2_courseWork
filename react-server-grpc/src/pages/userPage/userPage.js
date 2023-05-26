import { Box, Button, Stack, TextField, Typography } from "@mui/material";
import axios from "axios";
import React, { useCallback, useContext } from "react";
import { AuthContext } from "../../context/auth.context";
import OrderBlock from "../servicePages/ordersPage/orderBlock";

const UserPage = () => {
	const auth = useContext(AuthContext);
	const [user, setUser] = React.useState({
		name: "",
		email: "",
		newPassword: "",
		userType: "",
	});
	const [orders, setOrders] = React.useState([]);

	const getOrders = useCallback(async () => {
		await axios
			.get("http://localhost:8000/getorders", {
				headers: { Authorization: `Bearer ${auth.token}` },
			})
			.then((res) => {
				setOrders(res.data.orders);
			});
	}, []);
	const handleChange = (e) => {
		setUser({ ...user, [e.target.name]: e.target.value });
	};

	React.useEffect(() => {
		getOrders();
	}, [getOrders]);

	return (
		<>
			<Typography variant="h2">Страница пользователя</Typography>
			<Typography variant="h4">Личные данные пользователя</Typography>
			<Stack spacing={2}>
				<Box>
					Имя пользователя
					<TextField
						variant="outlined"
						name="name"
						fullWidth
						onChange={handleChange}
						defaultValue={user.name}
					/>
				</Box>
				<Box>
					Email пользователя
					<TextField
						type="email"
						name="email"
						variant="outlined"
						fullWidth
						onChange={handleChange}
						defaultValue={user.email}
					/>
				</Box>
				<Box>
					Новый пароль
					<TextField
						variant="outlined"
						fullWidth
						name="newPassword"
						onChange={handleChange}
					/>
				</Box>
				<Button>Сохранить</Button>
			</Stack>
			<Typography variant="h4">Заказы</Typography>
			{orders.map((order) => (
				<OrderBlock order={order} />
			))}
		</>
	);
};

export default UserPage;
