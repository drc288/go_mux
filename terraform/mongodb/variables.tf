variable mongodb {
  type        = object({
      private_key = string
      public_key  = string
  })
  description = "Object with private key and public key, this is using to connect a mongodb"
}

variable organization_id {
  type        = string
  default     = "5de47051c56c98420e72486d"
  description = "organization id in mongo db"
}
