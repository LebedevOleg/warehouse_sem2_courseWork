/* Триггер считает можно ли заказать данное колличество товаров на пользователя
 * если полльзователь физический и товара не хватает то возвращает ошибку
 * иначе все работает
 *  */
create or replace function AddOffer() returns trigger as $addOffer$
begin 
	if (select count from items where id = new.item_id) < new.item.count then 
		if (select type_id from users where id = (select user_id from orders where id = new.order_id)) != 1 then 
			delete from items_to_orders where order_id = new.order_id;
			delete from orders where id = order_id;
			raise notice 'на складе не хватает товара %', (select name from items where id = new.item_id);
			return null;
		end if;
		update orders set status = -1 where id = new.order_id;
	new.item_need_count = new.item_count - (select count from items where id = new.item_id);
	end if;
	update items set count = count - new.count where id = new.item_id;
	return new;
end
$addOffer$ language plpgsql;



create or replace trigger addOffer
	before insert on items_to_orders
	for each row execute procedure AddOffer()

	
create or replace function AddItem() returns trigger as $addItem$
begin 	
	if exists (select item_id from storage_to_items where item_id = new.item_id) then 
		update storage_to_items set item_count = item_count + new.item_count
			where storage_id = (select storage_id from deliveries where id = new.delivery_id) 
			and item_id = new.item_id;
	else
		insert into storage_to_items (storage_id, item_id,item_count)
			values ((select storage_id from deliveries where id = new.delivery_id),
				new.item_id, new.item_count);
	end if;
	return new;
end
$addItem$ language plpgsql;

create or replace trigger addItem
	before insert on providers_to_deliveries
	for each row execute procedure AddItem()
	
	
	
create or replace function UpdateItemInStorage() returns trigger as $updateItemInStorage$
begin 
	if (select count from items where id = new.item_id) < 0 then 
		new.item_count = new.item_count - (select item_need_count from items_to_orders 
			where item_need_count != 0 limit 1);
		update items_to_orders set item_need_count = 0 where 
			id = (select id from items_to_orders where item_need_count != 0 limit 1);		
	end if;
	update items set count = count - old.item_count + new.item_count where id = new.item_id;
end
$updateItemInStorage$ language plpgsql;

create or replace trigger updateItemInStorage
	after update or insert on storage_to_items
	for each row execute procedure UpdateItemInStorage()


	
