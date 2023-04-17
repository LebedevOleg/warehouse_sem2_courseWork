import { Button, Stack, Typography } from "@mui/material";
import React from "react";

const ServicePage = () => {
	return (
		<>
			<Typography variant="h2">Сервисная страница</Typography>
			<Stack spacing={2}>
				<Button
					sx={{ display: "block" }}
					onClick={() => {
						window.location = "/service/items";
					}}
				>
					<Typography variant="h5">Управление товарами</Typography>
				</Button>
				<Button
					sx={{ display: "block" }}
					onClick={() => {
						window.location = "/shop";
					}}
				>
					<Typography variant="h5">
						Управление поставщиками
					</Typography>
				</Button>
				<Button
					sx={{ display: "block" }}
					onClick={() => {
						window.location = "/shop";
					}}
				>
					<Typography variant="h5">Добавление поставки</Typography>
				</Button>
				<Button
					sx={{ display: "block" }}
					onClick={() => {
						window.location = "/shop";
					}}
				>
					<Typography variant="h5">Управление складами</Typography>
				</Button>
				<Button
					sx={{ display: "block" }}
					onClick={() => {
						window.location = "/shop";
					}}
				>
					<Typography variant="h5">
						Управление пользователями
					</Typography>
				</Button>
			</Stack>
		</>
	);
};

export default ServicePage;
