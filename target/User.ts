export interface IUser {
	userId : string;
	email : string;
	password : string;
	fullName ?: string;

}

interface UserProps {
	userId : string;
	email : string;
	password : string;
	fullName ?: string;

}

export default class User {
	state : IUser;
	constructor(props : IUser) {
		this.state = {...props};
	}


	set userId(userId:string) {
		this.state.userId=userId;
	}
	set email(email:string) {
		this.state.email=email;
	}
	set password(password:string) {
		this.state.password=password;
	}
	set fullName(fullName:string) {
		this.state.fullName=fullName;
	}



	get userId(): string {
		return this.state.userId;
	}
	get email(): string {
		return this.state.email;
	}
	get password(): string {
		return this.state.password;
	}
	get fullName(): string {
		return this.state.fullName;
	}

	toDTO() : IUser {
		return this.state;
	}
}
