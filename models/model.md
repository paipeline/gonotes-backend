| Documents Field | Type | Description |
|-----------------|----------------|-------------------------------------------|
| id | ID | Primary key, auto-generated |
| title | String | Document title |
| content | String | Document content (optional) |
| userId | String | Foreign key to users |
| parentDocument | ID (Documents) | Self-reference to parent document (optional)|
| isArchived | Boolean | Indicates if document is archived |
| isPublished | Boolean | Indicates if document is published |
| coverImage | String | URL or reference to cover image (optional) |
| icon | String | Icon identifier or reference (optional) |
| Indexes | |
|------------------------|---------------------------------------------------|
| by_user | On userId |
| by_user_parent | Composite index on (userId, parentDocument) |
| Relationships | |
|------------------------|---------------------------------------------------|
| Self-referential | parentDocument refers to another document's id |
| User association | userId links to an implied Users table |

Key Features:
Hierarchical structure (tree) of documents
2. User-specific document ownership
Archive and publish functionality
4. Support for cover images and icons
Efficient querying for user documents and child documents
This table structure supports the implemented mutations and queries, including document creation, archiving, restoring, updating, and various retrieval operations based on user and document hierarchy.