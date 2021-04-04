export default class UserMapper {
	public static toDomain(model : any) {
		return new User({
		userId: model.user_id,
		email: model.email,
		password: model.password,
		fullName: model.full_name,

})	}

	public static toPersistence(domainModel : User) {
		return {
		user_id: domainModel.userId,
		email: domainModel.email,
		password: domainModel.password,
		full_name: domainModel.fullName,

}	}

}
