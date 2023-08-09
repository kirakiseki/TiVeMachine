export type RoleType = '' | 'admin' | 'user'

export interface UserState {
  name?: string
  avatar?: string
  email?: string
  phone?: string
  registrationDate?: string
  accountId?: string
  role: RoleType
}
