import React from "react";
import { useAuth } from "../../../../hooks/auth.hook";
import axios from "axios";
import {
	Button,
	Dialog,
	DialogActions,
	DialogContent,
	DialogTitle,
	TextField,
} from "@mui/material";
import { Box } from "@mui/system";

const UserChangeModal = ({ user }) => {
	const [openModal, setOpenModal] = React.useState(false);
	const { token } = useAuth();
	const [newUser, setNewUser] = React.useState({
		id: user.id,
		name: user.name,
		email: user.email,
		type: user.type,
		role: user.role,
	});
	const handleOpenModal = () => {
		setOpenModal(true);
	};
	const handleCloseModal = () => {
		setOpenModal(false);
	};
	const handleChangeUserValue = (e) => {
		setNewUser({
			...user,
			[e.target.name]: e.target.value,
		});
	};
	const handleUpdateUser = async (e) => {
		await axios
			.post(`http://localhost:8000/updateuser`, newUser, {
				headers: {
					Authorization: `Bearer ${token}`,
				},
			})
			.then((res) => {
				setOpenModal(false);
			});
	};

	return (
		<div>
			<Button onClick={handleOpenModal}>Изменить</Button>
			<Dialog open={openModal} onClose={handleCloseModal} fullWidth>
				<DialogTitle>Изменить пользователя</DialogTitle>
				<DialogContent>
					<Box sx={{ display: "flex", alignItems: "flex-end" }}>
						<TextField
							defaultValue={newUser.name}
							name="name"
							margin="dense"
							id="name"
							onChange={handleChangeUserValue}
						/>
					</Box>
					<Box sx={{ display: "flex", alignItems: "flex-end" }}>
						<TextField
							defaultValue={newUser.email}
							name="email"
							margin="dense"
							id="email"
							onChange={handleChangeUserValue}
						/>
					</Box>
					<Box sx={{ display: "flex", alignItems: "flex-end" }}>
						<TextField
							defaultValue={newUser.type}
							name="type"
							margin="dense"
							id="type"
							onChange={handleChangeUserValue}
						/>
					</Box>
					<Box sx={{ display: "flex", alignItems: "flex-end" }}>
						<TextField
							defaultValue={newUser.role}
							name="role"
							margin="dense"
							id="role"
							onChange={handleChangeUserValue}
						/>
					</Box>
				</DialogContent>
				<DialogActions>
					<Button onClick={handleCloseModal}>Закрыть</Button>
					<Button onClick={handleUpdateUser}>Изменить</Button>
				</DialogActions>
			</Dialog>
		</div>
	);
};

export default UserChangeModal;
