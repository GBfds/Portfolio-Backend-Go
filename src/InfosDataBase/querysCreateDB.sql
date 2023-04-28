--Types

CREATE TYPE pedido AS(
  id_produto UUID,
  nome_produto TEXT,
  quantidade INT
);

--Tables
CREATE TABLE admin(
  id UUID DEFAULT uuid_generate_v4(),
  nome VARCHAR(200) NOT NULL,
  email TEXT UNIQUE NOT NULL,
  senha TEXT NOT NULL,
  cargo TEXT NOT NULL,
  
  CONSTRAINT id_admin_pkey PRIMARY KEY (id)
);

CREATE TABLE cliente(
  id UUID DEFAULT uuid_generate_v4(),
  nome VARCHAR(200) NOT NULL,
  email TEXT UNIQUE NOT NULL,
  senha TEXT NOT NULL,
  telefone CHAR(11) NOT NULL,
  
  CONSTRAINT id_cliente_pkey PRIMARY KEY (id)
);

CREATE TABLE endereco(
  id UUID DEFAULT uuid_generate_v4(),
  id_cliente UUID NOT NULL,
  numero VARCHAR(10),
  rua TEXT NOT NULL,
  bairro TEXT NOT NULL,
  cidade TEXT NOT NULL,
  
  CONSTRAINT id_endereco_pkey PRIMARY KEY (id),
  CONSTRAINT id_cliente_fkey FOREIGN KEY (id_cliente) REFERENCES cliente(id)
);


CREATE TABLE produto(
  id UUID DEFAULT uuid_generate_v4(),
  nome TEXT NOT NULL,
  tamanho TEXT NOT NULL,
  preco REAL NOT NULL,
  
  CONSTRAINT id_produto_pkey PRIMARY KEY (id)
);







