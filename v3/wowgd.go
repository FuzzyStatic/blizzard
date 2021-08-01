package blizzard

import (
	"context"
	"fmt"

	"github.com/FuzzyStatic/blizzard/wowgd"
	"github.com/FuzzyStatic/blizzard/wowsearch"
)

// WoWAchievementCategoriesIndex returns an index of achievement categories.
func (c *Client) WoWAchievementCategoriesIndex(ctx context.Context) (*wowgd.AchievementCategoriesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/achievement-category/index",
		c.GetStaticNamespace(),
		&wowgd.AchievementCategoriesIndex{},
	)
	return dat.(*wowgd.AchievementCategoriesIndex), header, err
}

// WoWAchievementCategory returns an achievement category by ID.
func (c *Client) WoWAchievementCategory(ctx context.Context, achievementCategoryID int) (*wowgd.AchievementCategory, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/achievement-category/%d", achievementCategoryID),
		c.GetStaticNamespace(),
		&wowgd.AchievementCategory{},
	)
	return dat.(*wowgd.AchievementCategory), header, err
}

// WoWAchievementIndex returns an index of achievements.
func (c *Client) WoWAchievementIndex(ctx context.Context) (*wowgd.AchievementIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/achievement/index",
		c.GetStaticNamespace(),
		&wowgd.AchievementIndex{},
	)
	return dat.(*wowgd.AchievementIndex), header, err
}

// WoWAchievement returns an achievement category by ID.
func (c *Client) WoWAchievement(ctx context.Context, achievementID int) (*wowgd.Achievement, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/achievement/%d", achievementID),
		c.GetStaticNamespace(),
		&wowgd.Achievement{},
	)
	return dat.(*wowgd.Achievement), header, err
}

// WoWAchievementMedia returns media for an achievement by ID.
func (c *Client) WoWAchievementMedia(ctx context.Context, achievementID int) (*wowgd.AchievementMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/achievement/%d", achievementID),
		c.GetStaticNamespace(),
		&wowgd.AchievementMedia{},
	)
	return dat.(*wowgd.AchievementMedia), header, err
}

// WoWAuctions returns all active auctions for a connected realm. See the Connected Realm API for information about retrieving a list of connected realm IDs.
// Auction house data updates at a set interval. The value was initially set at 1 hour; however, it might change over time without notice.
// Depending on the number of active auctions on the specified connected realm, the response from this endpoint may be rather large, sometimes exceeding 10 MB.
func (c *Client) WoWAuctions(ctx context.Context, connectedRealmID int) (*wowgd.AuctionHouse, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/auctions", connectedRealmID),
		c.GetDynamicNamespace(),
		&wowgd.AuctionHouse{},
	)
	return dat.(*wowgd.AuctionHouse), header, err
}

// WoWAzeriteEssenceIndex returns an index of azerite essences.
func (c *Client) WoWAzeriteEssenceIndex(ctx context.Context) (*wowgd.AzeriteEssenceIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/azerite-essence/index",
		c.GetStaticNamespace(),
		&wowgd.AzeriteEssenceIndex{},
	)
	return dat.(*wowgd.AzeriteEssenceIndex), header, err
}

// WoWAzeriteEssence returns an azerite essence by ID.
func (c *Client) WoWAzeriteEssence(ctx context.Context, azeriteEssenceID int) (*wowgd.AzeriteEssence, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/azerite-essence/%d", azeriteEssenceID), c.GetStaticNamespace(),
		&wowgd.AzeriteEssence{},
	)
	return dat.(*wowgd.AzeriteEssence), header, err
}

// WoWAzeriteEssenceMedia returns media for an azerite essence by ID.
func (c *Client) WoWAzeriteEssenceMedia(ctx context.Context, azeriteEssenceID int) (*wowgd.AzeriteEssenceMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/azerite-essence/%d", azeriteEssenceID),
		c.GetStaticNamespace(),
		&wowgd.AzeriteEssenceMedia{},
	)
	return dat.(*wowgd.AzeriteEssenceMedia), header, err
}

// WoWConnectedRealmsIndex returns an index of connected realms
func (c *Client) WoWConnectedRealmsIndex(ctx context.Context) (*wowgd.ConnectedRealmsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/connected-realm/index",
		c.GetDynamicNamespace(),
		&wowgd.ConnectedRealmsIndex{},
	)
	return dat.(*wowgd.ConnectedRealmsIndex), header, err
}

// WoWConnectedRealm returns a single connected realm by ID
func (c *Client) WoWConnectedRealm(ctx context.Context, connectedRealmID int) (*wowgd.ConnectedRealm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d", connectedRealmID),
		c.GetDynamicNamespace(),
		&wowgd.ConnectedRealm{},
	)
	return dat.(*wowgd.ConnectedRealm), header, err
}

// WoWCovenantsIndex returns an index of covenants.
func (c *Client) WoWCovenantsIndex(ctx context.Context) (*wowgd.CovenantsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/covenant/index",
		c.GetStaticNamespace(),
		&wowgd.CovenantsIndex{},
	)
	return dat.(*wowgd.CovenantsIndex), header, err
}

// WoWConnectedRealmSearch searches for connected realms
func (c *Client) WoWConnectedRealmSearch(ctx context.Context, opts ...wowsearch.Opt) (*wowgd.ConnectedRealmsSearch, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/search/connected-realm%s", buildSearchParams(opts...)),
		c.GetDynamicClassicNamespace(),
		&wowgd.ConnectedRealmsSearch{},
	)
	return dat.(*wowgd.ConnectedRealmsSearch), header, err
}

// WoWCovenant returns a covenant by ID.
func (c *Client) WoWCovenant(ctx context.Context, covenantID int) (*wowgd.Covenant, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/covenant/%d", covenantID),
		c.GetStaticNamespace(),
		&wowgd.Covenant{},
	)
	return dat.(*wowgd.Covenant), header, err
}

// WoWCovenantMedia returns media for a covenant by ID.
func (c *Client) WoWCovenantMedia(ctx context.Context, covenantID int) (*wowgd.CovenantMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/covenant/%d", covenantID),
		c.GetStaticNamespace(),
		&wowgd.CovenantMedia{},
	)
	return dat.(*wowgd.CovenantMedia), header, err
}

// WoWCovenantSoulbindsIndex returns an index of soulbinds.
func (c *Client) WoWCovenantSoulbindsIndex(ctx context.Context) (*wowgd.CovenantSoulbindsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/covenant/soulbind/index",
		c.GetStaticNamespace(),
		&wowgd.CovenantSoulbindsIndex{},
	)
	return dat.(*wowgd.CovenantSoulbindsIndex), header, err
}

// WoWCovenantSoulbind returns a soulbind by ID.
func (c *Client) WoWCovenantSoulbind(ctx context.Context, soulbindID int) (*wowgd.CovenantSoulbind, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/covenant/soulbind/%d", soulbindID),
		c.GetStaticNamespace(),
		&wowgd.CovenantSoulbind{},
	)
	return dat.(*wowgd.CovenantSoulbind), header, err
}

// WoWCovenantConduitsIndex returns an index of conduits.
func (c *Client) WoWCovenantConduitsIndex(ctx context.Context) (*wowgd.CovenantConduitsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/covenant/conduit/index",
		c.GetStaticNamespace(),
		&wowgd.CovenantConduitsIndex{},
	)
	return dat.(*wowgd.CovenantConduitsIndex), header, err
}

// WoWCovenantConduit returns a conduit by ID.
func (c *Client) WoWCovenantConduit(ctx context.Context, conduitID int) (*wowgd.CovenantConduit, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/covenant/conduit/%d", conduitID),
		c.GetStaticNamespace(),
		&wowgd.CovenantConduit{},
	)
	return dat.(*wowgd.CovenantConduit), header, err
}

// WoWCreatureFamiliesIndex returns an index of creature families.
func (c *Client) WoWCreatureFamiliesIndex(ctx context.Context) (*wowgd.CreatureFamiliesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-family/index",
		c.GetStaticNamespace(),
		&wowgd.CreatureFamiliesIndex{},
	)
	return dat.(*wowgd.CreatureFamiliesIndex), header, err
}

// WoWCreatureFamily returns a creature family by ID.
func (c *Client) WoWCreatureFamily(ctx context.Context, creatureFamilyID int) (*wowgd.CreatureFamily, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-family/%d", creatureFamilyID),
		c.GetStaticNamespace(),
		&wowgd.CreatureFamily{},
	)
	return dat.(*wowgd.CreatureFamily), header, err
}

// WoWCreatureTypesIndex returns an index of creature types.
func (c *Client) WoWCreatureTypesIndex(ctx context.Context) (*wowgd.CreatureTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-type/index",
		c.GetStaticNamespace(),
		&wowgd.CreatureTypesIndex{},
	)
	return dat.(*wowgd.CreatureTypesIndex), header, err
}

// WoWCreatureType returns a creature type by ID.
func (c *Client) WoWCreatureType(ctx context.Context, creatureTypeID int) (*wowgd.CreatureType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-type/%d", creatureTypeID),
		c.GetStaticNamespace(),
		&wowgd.CreatureType{},
	)
	return dat.(*wowgd.CreatureType), header, err
}

// WoWCreature returns a creature type by ID.
func (c *Client) WoWCreature(ctx context.Context, creatureID int) (*wowgd.Creature, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature/%d", creatureID),
		c.GetStaticNamespace(),
		&wowgd.Creature{},
	)
	return dat.(*wowgd.Creature), header, err
}

// WoWCreatureDisplayMedia returns media for a creature display by ID.
func (c *Client) WoWCreatureDisplayMedia(ctx context.Context, creatureDisplayID int) (*wowgd.CreatureDisplayMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-display/%d", creatureDisplayID),
		c.GetStaticNamespace(),
		&wowgd.CreatureDisplayMedia{},
	)
	return dat.(*wowgd.CreatureDisplayMedia), header, err
}

// WoWCreatureFamilyMedia returns media for a creature family by ID.
func (c *Client) WoWCreatureFamilyMedia(ctx context.Context, creatureFamilyID int) (*wowgd.CreatureFamilyMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-family/%d", creatureFamilyID),
		c.GetStaticNamespace(),
		&wowgd.CreatureFamilyMedia{},
	)
	return dat.(*wowgd.CreatureFamilyMedia), header, err
}

// WoWGuildCrestComponentsIndex returns an index of guild crest media.
func (c *Client) WoWGuildCrestComponentsIndex(ctx context.Context) (*wowgd.GuildCrestComponentsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/guild-crest/index",
		c.GetStaticNamespace(),
		&wowgd.GuildCrestComponentsIndex{},
	)
	return dat.(*wowgd.GuildCrestComponentsIndex), header, err
}

// WoWGuildCrestBorderMedia returns media for a guild crest border by ID.
func (c *Client) WoWGuildCrestBorderMedia(ctx context.Context, borderID int) (*wowgd.GuildCrestBorderMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/border/%d", borderID),
		c.GetStaticNamespace(),
		&wowgd.GuildCrestBorderMedia{},
	)
	return dat.(*wowgd.GuildCrestBorderMedia), header, err
}

// WoWGuildCrestEmblemMedia returns media for a guild crest emblem by ID.
func (c *Client) WoWGuildCrestEmblemMedia(ctx context.Context, emblemID int) (*wowgd.GuildCrestEmblemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/emblem/%d", emblemID),
		c.GetStaticNamespace(),
		&wowgd.GuildCrestEmblemMedia{},
	)
	return dat.(*wowgd.GuildCrestEmblemMedia), header, err
}

// WoWItemClassesIndex returns an index of item classes.
func (c *Client) WoWItemClassesIndex(ctx context.Context) (*wowgd.ItemClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/item-class/index",
		c.GetStaticNamespace(),
		&wowgd.ItemClassesIndex{},
	)
	return dat.(*wowgd.ItemClassesIndex), header, err
}

// WoWItemClass returns an item class by ID.
func (c *Client) WoWItemClass(ctx context.Context, itemClassID int) (*wowgd.ItemClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d", itemClassID),
		c.GetStaticNamespace(),
		&wowgd.ItemClass{},
	)
	return dat.(*wowgd.ItemClass), header, err
}

// WoWItemSetsIndex returns an index of item sets.
func (c *Client) WoWItemSetsIndex(ctx context.Context) (*wowgd.ItemSetsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/item-set/index",
		c.GetStaticNamespace(),
		&wowgd.ItemSetsIndex{},
	)
	return dat.(*wowgd.ItemSetsIndex), header, err
}

// WoWItemSet returns an item set by ID.
func (c *Client) WoWItemSet(ctx context.Context, itemSetID int) (*wowgd.ItemSet, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-set/%d", itemSetID),
		c.GetStaticNamespace(),
		&wowgd.ItemSet{},
	)
	return dat.(*wowgd.ItemSet), header, err
}

// WoWItemSubclass returns an item subclass by ID.
func (c *Client) WoWItemSubclass(ctx context.Context, itemClassID, itemSubclassID int) (*wowgd.ItemSubclass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d/item-subclass/%d", itemClassID, itemSubclassID),
		c.GetStaticNamespace(),
		&wowgd.ItemSubclass{},
	)
	return dat.(*wowgd.ItemSubclass), header, err
}

// WoWItem returns an item by ID.
func (c *Client) WoWItem(ctx context.Context, itemID int) (*wowgd.Item, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item/%d", itemID),
		c.GetStaticNamespace(),
		&wowgd.Item{},
	)
	return dat.(*wowgd.Item), header, err
}

// WoWItemMedia returns media for an item by ID.
func (c *Client) WoWItemMedia(ctx context.Context, itemID int) (*wowgd.ItemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/item/%d", itemID),
		c.GetStaticNamespace(),
		&wowgd.ItemMedia{},
	)
	return dat.(*wowgd.ItemMedia), header, err
}

// WoWJournalExpansionsIndex returns an index of journal expansions.
func (c *Client) WoWJournalExpansionsIndex(ctx context.Context) (*wowgd.JournalExpansionsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/journal-expansion/index",
		c.GetStaticNamespace(),
		&wowgd.JournalExpansionsIndex{},
	)
	return dat.(*wowgd.JournalExpansionsIndex), header, err
}

// WoWJournalExpansion returns a journal expansion by ID.
func (c *Client) WoWJournalExpansion(ctx context.Context, journalExpansionID int) (*wowgd.JournalExpansion, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/journal-expansion/%d", journalExpansionID),
		c.GetStaticNamespace(),
		&wowgd.JournalExpansion{},
	)
	return dat.(*wowgd.JournalExpansion), header, err
}

// WoWJournalEncountersIndex returns an index of journal encounters.
func (c *Client) WoWJournalEncountersIndex(ctx context.Context) (*wowgd.JournalEncountersIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/journal-encounter/index",
		c.GetStaticNamespace(),
		&wowgd.JournalEncountersIndex{},
	)
	return dat.(*wowgd.JournalEncountersIndex), header, err
}

// WoWJournalEncounter returns a journal expansion by ID.
func (c *Client) WoWJournalEncounter(ctx context.Context, journalEncounterID int) (*wowgd.JournalEncounter, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/journal-encounter/%d", journalEncounterID),
		c.GetStaticNamespace(),
		&wowgd.JournalEncounter{},
	)
	return dat.(*wowgd.JournalEncounter), header, err
}

// WoWJournalInstancesIndex returns an index of journal instances.
func (c *Client) WoWJournalInstancesIndex(ctx context.Context) (*wowgd.JournalInstancesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/journal-instance/index",
		c.GetStaticNamespace(),
		&wowgd.JournalInstancesIndex{},
	)
	return dat.(*wowgd.JournalInstancesIndex), header, err
}

// WoWJournalInstance returns a journal instance.
func (c *Client) WoWJournalInstance(ctx context.Context, journalInstanceID int) (*wowgd.JournalInstance, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/journal-instance/%d", journalInstanceID),
		c.GetStaticNamespace(),
		&wowgd.JournalInstance{},
	)
	return dat.(*wowgd.JournalInstance), header, err
}

// WoWJournalInstanceMedia returns media for a journal instance by ID.
func (c *Client) WoWJournalInstanceMedia(ctx context.Context, journalInstanceID int) (*wowgd.JournalInstanceMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/journal-instance/%d", journalInstanceID),
		c.GetStaticNamespace(),
		&wowgd.JournalInstanceMedia{},
	)
	return dat.(*wowgd.JournalInstanceMedia), header, err
}

// WoWModifiedCraftingIndex returns the parent index for Modified Crafting.
func (c *Client) WoWModifiedCraftingIndex(ctx context.Context) (*wowgd.ModifiedCraftingIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/modified-crafting/index",
		c.GetStaticNamespace(),
		&wowgd.ModifiedCraftingIndex{},
	)
	return dat.(*wowgd.ModifiedCraftingIndex), header, err
}

// WoWModifiedCraftingCategoryIndex returns the index of Modified Crafting categories.
func (c *Client) WoWModifiedCraftingCategoryIndex(ctx context.Context) (*wowgd.ModifiedCraftingCategoryIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/modified-crafting/category/index",
		c.GetStaticNamespace(),
		&wowgd.ModifiedCraftingCategoryIndex{},
	)
	return dat.(*wowgd.ModifiedCraftingCategoryIndex), header, err
}

// WoWModifiedCraftingCategory returns a Modified Crafting category by ID.
func (c *Client) WoWModifiedCraftingCategory(ctx context.Context, categoryID int) (*wowgd.ModifiedCraftingCategory, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/modified-crafting/category/%d", categoryID),
		c.GetStaticNamespace(),
		&wowgd.ModifiedCraftingCategory{},
	)
	return dat.(*wowgd.ModifiedCraftingCategory), header, err
}

// WoWModifiedCraftingReagentSlotTypeIndex returns the index of Modified Crafting reagent slot types.
func (c *Client) WoWModifiedCraftingReagentSlotTypeIndex(ctx context.Context) (*wowgd.ModifiedCraftingReagentSlotTypeIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/modified-crafting/reagent-slot-type/index",
		c.GetStaticNamespace(),
		&wowgd.ModifiedCraftingReagentSlotTypeIndex{},
	)
	return dat.(*wowgd.ModifiedCraftingReagentSlotTypeIndex), header, err
}

// WoWModifiedCraftingReagentSlotType returns a Modified Crafting reagent slot type by ID.
func (c *Client) WoWModifiedCraftingReagentSlotType(ctx context.Context,
	slotTypeID int) (*wowgd.ModifiedCraftingReagentSlotType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/modified-crafting/reagent-slot-type/%d", slotTypeID),
		c.GetStaticNamespace(),
		&wowgd.ModifiedCraftingReagentSlotType{},
	)
	return dat.(*wowgd.ModifiedCraftingReagentSlotType), header, err
}

// WoWMountIndex returns an index of mounts.
func (c *Client) WoWMountIndex(ctx context.Context) (*wowgd.MountIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/mount/index",
		c.GetStaticNamespace(),
		&wowgd.MountIndex{},
	)
	return dat.(*wowgd.MountIndex), header, err
}

// WoWMount returns a mount by ID.
func (c *Client) WoWMount(ctx context.Context, mountID int) (*wowgd.Mount, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/mount/%d", mountID),
		c.GetStaticNamespace(),
		&wowgd.Mount{},
	)
	return dat.(*wowgd.Mount), header, err
}

// WoWMythicKeystoneAffixIndex returns an index of Keystone affixes
func (c *Client) WoWMythicKeystoneAffixIndex(ctx context.Context) (*wowgd.MythicKeystoneAffixIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/keystone-affix/index",
		c.GetStaticNamespace(),
		&wowgd.MythicKeystoneAffixIndex{},
	)
	return dat.(*wowgd.MythicKeystoneAffixIndex), header, err
}

// WoWMythicKeystoneAffix returns a single connected realm by ID
func (c *Client) WoWMythicKeystoneAffix(ctx context.Context, keystoneAffixID int) (*wowgd.MythicKeystoneAffix, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/keystone-affix/%d", keystoneAffixID),
		c.GetStaticNamespace(),
		&wowgd.MythicKeystoneAffix{},
	)
	return dat.(*wowgd.MythicKeystoneAffix), header, err
}

// WoWMythicKeystoneAffixMedia returns media for a mythic keystone affix by ID.
func (c *Client) WoWMythicKeystoneAffixMedia(ctx context.Context, keystoneAffixID int) (*wowgd.MythicKeystoneAffixMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/keystone-affix/%d", keystoneAffixID),
		c.GetStaticNamespace(),
		&wowgd.MythicKeystoneAffixMedia{},
	)
	return dat.(*wowgd.MythicKeystoneAffixMedia), header, err
}

// WoWMythicKeystoneDungeonIndex returns an index of Mythic Keystone dungeons
func (c *Client) WoWMythicKeystoneDungeonIndex(ctx context.Context) (*wowgd.MythicKeystoneDungeonIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/mythic-keystone/dungeon/index",
		c.GetDynamicNamespace(),
		&wowgd.MythicKeystoneDungeonIndex{},
	)
	return dat.(*wowgd.MythicKeystoneDungeonIndex), header, err
}

// WoWMythicKeystoneDungeon returns a Mythic Keystone dungeon by ID
func (c *Client) WoWMythicKeystoneDungeon(ctx context.Context, dungeonID int) (*wowgd.MythicKeystoneDungeon, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/mythic-keystone/dungeon/%d", dungeonID),
		c.GetDynamicNamespace(),
		&wowgd.MythicKeystoneDungeon{},
	)
	return dat.(*wowgd.MythicKeystoneDungeon), header, err
}

// WoWMythicKeystoneIndex returns n index of links to other documents related to Mythic Keystone dungeons
func (c *Client) WoWMythicKeystoneIndex(ctx context.Context) (*wowgd.MythicKeystoneIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/mythic-keystone/index",
		c.GetDynamicNamespace(),
		&wowgd.MythicKeystoneIndex{},
	)
	return dat.(*wowgd.MythicKeystoneIndex), header, err
}

// WoWMythicKeystonePeriodIndex returns an index of Mythic Keystone periods
func (c *Client) WoWMythicKeystonePeriodIndex(ctx context.Context) (*wowgd.MythicKeystonePeriodIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/mythic-keystone/period/index",
		c.GetDynamicNamespace(),
		&wowgd.MythicKeystonePeriodIndex{},
	)
	return dat.(*wowgd.MythicKeystonePeriodIndex), header, err
}

// WoWMythicKeystonePeriod returns a Mythic Keystone period by ID
func (c *Client) WoWMythicKeystonePeriod(ctx context.Context, periodID int) (*wowgd.MythicKeystonePeriod, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/mythic-keystone/period/%d", periodID),
		c.GetDynamicNamespace(),
		&wowgd.MythicKeystonePeriod{},
	)
	return dat.(*wowgd.MythicKeystonePeriod), header, err
}

// WoWMythicKeystoneSeasonIndex returns an index of Mythic Keystone seasons
func (c *Client) WoWMythicKeystoneSeasonIndex(ctx context.Context) (*wowgd.MythicKeystoneSeasonIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/mythic-keystone/season/index",
		c.GetDynamicNamespace(),
		&wowgd.MythicKeystoneSeasonIndex{},
	)
	return dat.(*wowgd.MythicKeystoneSeasonIndex), header, err
}

// WoWMythicKeystoneSeason returns a Mythic Keystone season by ID
func (c *Client) WoWMythicKeystoneSeason(ctx context.Context, seasonID int) (*wowgd.MythicKeystoneSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/mythic-keystone/season/%d", seasonID),
		c.GetDynamicNamespace(),
		&wowgd.MythicKeystoneSeason{},
	)
	return dat.(*wowgd.MythicKeystoneSeason), header, err
}

// WoWMythicKeystoneLeaderboardIndex returns an index of Mythic Keystone Leaderboard dungeon instances for a connected realm
func (c *Client) WoWMythicKeystoneLeaderboardIndex(ctx context.Context,
	connectedRealmID int) (*wowgd.MythicKeystoneLeaderboardIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/mythic-leaderboard/index", connectedRealmID),
		c.GetDynamicNamespace(),
		&wowgd.MythicKeystoneLeaderboardIndex{},
	)
	return dat.(*wowgd.MythicKeystoneLeaderboardIndex), header, err
}

// WoWMythicKeystoneLeaderboard returns a weekly Mythic Keystone Leaderboard by period
func (c *Client) WoWMythicKeystoneLeaderboard(ctx context.Context,
	connectedRealmID, dungeonID, period int) (*wowgd.MythicKeystoneLeaderboard, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/mythic-leaderboard/%d/period/%d",
			connectedRealmID, dungeonID, period),
		c.GetDynamicNamespace(),
		&wowgd.MythicKeystoneLeaderboard{},
	)
	return dat.(*wowgd.MythicKeystoneLeaderboard), header, err
}

// WoWMythicRaidLeaderboard returns the leaderboard for a given raid and faction
func (c *Client) WoWMythicRaidLeaderboard(ctx context.Context, raid, faction string) (*wowgd.MythicRaidLeaderboard, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/leaderboard/hall-of-fame/%s/%s", raid, faction),
		c.GetDynamicNamespace(),
		&wowgd.MythicRaidLeaderboard{},
	)
	return dat.(*wowgd.MythicRaidLeaderboard), header, err
}

// WoWPetIndex returns an index of pets.
func (c *Client) WoWPetIndex(ctx context.Context) (*wowgd.PetIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pet/index",
		c.GetStaticNamespace(),
		&wowgd.PetIndex{},
	)
	return dat.(*wowgd.PetIndex), header, err
}

// WoWPet returns a pet by ID.
func (c *Client) WoWPet(ctx context.Context, petID int) (*wowgd.Pet, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pet/%d", petID),
		c.GetStaticNamespace(),
		&wowgd.Pet{},
	)
	return dat.(*wowgd.Pet), header, err
}

// WoWPetMedia returns media information for a pet ID.
func (c *Client) WoWPetMedia(ctx context.Context, petID int) (*wowgd.PetMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/pet/%d", petID),
		c.GetStaticNamespace(),
		&wowgd.PetMedia{},
	)
	return dat.(*wowgd.PetMedia), header, err
}

// WoWPetAbilityIndex returns an index of pet abilities.
func (c *Client) WoWPetAbilityIndex(ctx context.Context) (*wowgd.PetAbilityIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pet-ability/index",
		c.GetStaticNamespace(),
		&wowgd.PetAbilityIndex{},
	)
	return dat.(*wowgd.PetAbilityIndex), header, err
}

// WoWPetAbility returns a pet ability by ID.
func (c *Client) WoWPetAbility(ctx context.Context, petAbilityID int) (*wowgd.PetAbility, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pet-ability/%d", petAbilityID),
		c.GetStaticNamespace(),
		&wowgd.PetAbility{},
	)
	return dat.(*wowgd.PetAbility), header, err
}

// WoWPetAbilityMedia returns media information for a pet ID.
func (c *Client) WoWPetAbilityMedia(ctx context.Context, petAbilityID int) (*wowgd.PetMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/pet-ability/%d", petAbilityID),
		c.GetStaticNamespace(),
		&wowgd.PetMedia{},
	)
	return dat.(*wowgd.PetMedia), header, err
}

// WoWPlayableClassesIndex returns an index of playable classes
func (c *Client) WoWPlayableClassesIndex(ctx context.Context) (*wowgd.PlayableClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-class/index",
		c.GetStaticNamespace(),
		&wowgd.PlayableClassesIndex{},
	)
	return dat.(*wowgd.PlayableClassesIndex), header, err
}

// WoWPlayableClass returns a playable class by ID
func (c *Client) WoWPlayableClass(ctx context.Context, classID int) (*wowgd.PlayableClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-class/%d", classID),
		c.GetStaticNamespace(),
		&wowgd.PlayableClass{},
	)
	return dat.(*wowgd.PlayableClass), header, err
}

// WoWPlayableClassMedia returns media for a playable class by ID.
func (c *Client) WoWPlayableClassMedia(ctx context.Context, classID int) (*wowgd.PlayableClassMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/playable-class/%d", classID),
		c.GetStaticNamespace(),
		&wowgd.PlayableClassMedia{},
	)
	return dat.(*wowgd.PlayableClassMedia), header, err
}

// WoWPlayableClassPvPTalentSlots returns the PvP talent slots for a playable class by ID
func (c *Client) WoWPlayableClassPvPTalentSlots(ctx context.Context, classID int) (*wowgd.PlayableClassPvPTalentSlots, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-class/%d/pvp-talent-slots", classID),
		c.GetStaticNamespace(),
		&wowgd.PlayableClassPvPTalentSlots{},
	)
	return dat.(*wowgd.PlayableClassPvPTalentSlots), header, err
}

// WoWPlayableRacesIndex returns an index of races.
func (c *Client) WoWPlayableRacesIndex(ctx context.Context) (*wowgd.PlayableRacesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-race/index",
		c.GetStaticNamespace(),
		&wowgd.PlayableRacesIndex{},
	)
	return dat.(*wowgd.PlayableRacesIndex), header, err
}

// WoWPlayableRace returns a race by ID.
func (c *Client) WoWPlayableRace(ctx context.Context, raceID int) (*wowgd.PlayableRace, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-race/%d", raceID),
		c.GetStaticNamespace(),
		&wowgd.PlayableRace{},
	)
	return dat.(*wowgd.PlayableRace), header, err
}

// WoWPlayableSpecializationIndex returns an index of playable specializations.
func (c *Client) WoWPlayableSpecializationIndex(ctx context.Context) (*wowgd.PlayableSpecializationIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-specialization/index",
		c.GetStaticNamespace(),
		&wowgd.PlayableSpecializationIndex{},
	)
	return dat.(*wowgd.PlayableSpecializationIndex), header, err
}

// WoWPlayableSpecialization returns a playable specialization by ID.
func (c *Client) WoWPlayableSpecialization(ctx context.Context, specID int) (*wowgd.PlayableSpecialization, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-specialization/%d", specID),
		c.GetStaticNamespace(),
		&wowgd.PlayableSpecialization{},
	)
	return dat.(*wowgd.PlayableSpecialization), header, err
}

// WoWPlayableSpecializationMedia returns media for a playable specialization by ID.
func (c *Client) WoWPlayableSpecializationMedia(ctx context.Context, specID int) (*wowgd.PlayableSpecializationMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/playable-specialization/%d", specID),
		c.GetStaticNamespace(),
		&wowgd.PlayableSpecializationMedia{},
	)
	return dat.(*wowgd.PlayableSpecializationMedia), header, err
}

// WoWPowerTypesIndex returns an index of power types.
func (c *Client) WoWPowerTypesIndex(ctx context.Context) (*wowgd.PowerTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/power-type/index",
		c.GetStaticNamespace(),
		&wowgd.PowerTypesIndex{},
	)
	return dat.(*wowgd.PowerTypesIndex), header, err
}

// WoWPowerType returns a power type by ID.
func (c *Client) WoWPowerType(ctx context.Context, powerTypeID int) (*wowgd.PowerType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/power-type/%d", powerTypeID),
		c.GetStaticNamespace(),
		&wowgd.PowerType{},
	)
	return dat.(*wowgd.PowerType), header, err
}

// WoWProfessionsIndex returns an index of professions.
func (c *Client) WoWProfessionsIndex(ctx context.Context) (*wowgd.ProfessionsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/profession/index",
		c.GetStaticNamespace(),
		&wowgd.ProfessionsIndex{},
	)
	return dat.(*wowgd.ProfessionsIndex), header, err
}

// WoWProfession returns a profession by ID.
func (c *Client) WoWProfession(ctx context.Context, professionID int) (*wowgd.Profession, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/profession/%d", professionID),
		c.GetStaticNamespace(),
		&wowgd.Profession{},
	)
	return dat.(*wowgd.Profession), header, err
}

// WoWProfessionMedia returns media for a profession by ID.
func (c *Client) WoWProfessionMedia(ctx context.Context, professionID int) (*wowgd.ProfessionMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/profession/%d", professionID),
		c.GetStaticNamespace(),
		&wowgd.ProfessionMedia{},
	)
	return dat.(*wowgd.ProfessionMedia), header, err
}

// WoWProfessionSkillTier returns a skill tier for a profession by ID.
func (c *Client) WoWProfessionSkillTier(ctx context.Context, professionID, skillTierID int) (*wowgd.ProfessionSkillTier, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/profession/%d/skill-tier/%d", professionID, skillTierID),
		c.GetStaticNamespace(),
		&wowgd.ProfessionSkillTier{},
	)
	return dat.(*wowgd.ProfessionSkillTier), header, err
}

// WoWRecipe returns a recipe by ID.
func (c *Client) WoWRecipe(ctx context.Context, recipeID int) (*wowgd.Recipe, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/recipe/%d", recipeID),
		c.GetStaticNamespace(),
		&wowgd.Recipe{},
	)
	return dat.(*wowgd.Recipe), header, err
}

// WoWRecipeMedia returns media for a recipe by ID.
func (c *Client) WoWRecipeMedia(ctx context.Context, recipeID int) (*wowgd.RecipeMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/recipe/%d", recipeID),
		c.GetStaticNamespace(),
		&wowgd.RecipeMedia{},
	)
	return dat.(*wowgd.RecipeMedia), header, err
}

// WoWPvPSeasonIndex returns an index of PvP seasons.
func (c *Client) WoWPvPSeasonIndex(ctx context.Context) (*wowgd.PvPSeasonIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pvp-season/index",
		c.GetDynamicNamespace(),
		&wowgd.PvPSeasonIndex{},
	)
	return dat.(*wowgd.PvPSeasonIndex), header, err
}

// WoWPvPSeason returns a PvP season by ID.
func (c *Client) WoWPvPSeason(ctx context.Context, pvpSeasonID int) (*wowgd.PvPSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-season/%d", pvpSeasonID),
		c.GetDynamicNamespace(),
		&wowgd.PvPSeason{},
	)
	return dat.(*wowgd.PvPSeason), header, err
}

// WoWPvPLeaderboardsIndex returns an index of PvP leaderboards for a PvP season.
func (c *Client) WoWPvPLeaderboardsIndex(ctx context.Context, pvpSeasonID int) (*wowgd.PvPLeaderboardsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-season/%d/pvp-leaderboard/index", pvpSeasonID),
		c.GetDynamicNamespace(),
		&wowgd.PvPLeaderboardsIndex{},
	)
	return dat.(*wowgd.PvPLeaderboardsIndex), header, err
}

// WoWPvPLeaderboard returns the PvP leaderboard of a specific PvP bracket for a PvP season.
func (c *Client) WoWPvPLeaderboard(ctx context.Context, pvpSeasonID int, pvpBracket wowgd.Bracket) (*wowgd.PvPLeaderboard, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-season/%d/pvp-leaderboard/%s", pvpSeasonID, pvpBracket),
		c.GetDynamicNamespace(),
		&wowgd.PvPLeaderboard{},
	)
	return dat.(*wowgd.PvPLeaderboard), header, err
}

// WoWPvPRewardsIndex returns an index of PvP rewards for a PvP season.
func (c *Client) WoWPvPRewardsIndex(ctx context.Context, pvpSeasonID int) (*wowgd.PvPRewardsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-season/%d/pvp-reward/index", pvpSeasonID),
		c.GetDynamicNamespace(),
		&wowgd.PvPRewardsIndex{},
	)
	return dat.(*wowgd.PvPRewardsIndex), header, err
}

// WoWPvPTierMedia returns media for a PvP tier by ID.
func (c *Client) WoWPvPTierMedia(ctx context.Context, pvpTierID int) (*wowgd.PvPTierMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/pvp-tier/%d", pvpTierID),
		c.GetStaticNamespace(),
		&wowgd.PvPTierMedia{},
	)
	return dat.(*wowgd.PvPTierMedia), header, err
}

// WoWPvPTiersIndex returns an index of PvP tiers.
func (c *Client) WoWPvPTiersIndex(ctx context.Context) (*wowgd.PvPTiersIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pvp-tier/index",
		c.GetStaticNamespace(),
		&wowgd.PvPTiersIndex{},
	)
	return dat.(*wowgd.PvPTiersIndex), header, err
}

// WoWPvPTier returns a PvP tier by ID.
func (c *Client) WoWPvPTier(ctx context.Context, pvpTierID int) (*wowgd.PvPTier, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-tier/%d", pvpTierID),
		c.GetStaticNamespace(),
		&wowgd.PvPTier{},
	)
	return dat.(*wowgd.PvPTier), header, err
}

// WoWQuestsIndex returns the parent index for quests.
func (c *Client) WoWQuestsIndex(ctx context.Context) (*wowgd.QuestsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/quest/index",
		c.GetStaticNamespace(),
		&wowgd.QuestsIndex{},
	)
	return dat.(*wowgd.QuestsIndex), header, err
}

// WoWQuest returns a quest by ID.
func (c *Client) WoWQuest(ctx context.Context, questID int) (*wowgd.Quest, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/quest/%d", questID),
		c.GetStaticNamespace(),
		&wowgd.Quest{},
	)
	return dat.(*wowgd.Quest), header, err
}

// WoWQuestCategoriesIndex returns an index of quest categories (such as quests for a specific class, profession, or storyline).
func (c *Client) WoWQuestCategoriesIndex(ctx context.Context) (*wowgd.QuestCategoriesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/quest/category/index",
		c.GetStaticNamespace(),
		&wowgd.QuestCategoriesIndex{},
	)
	return dat.(*wowgd.QuestCategoriesIndex), header, err
}

// WoWQuestCategory returns a quest category by ID.
func (c *Client) WoWQuestCategory(ctx context.Context, questCategoryID int) (*wowgd.QuestCategory, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/quest/category/%d", questCategoryID),
		c.GetStaticNamespace(),
		&wowgd.QuestCategory{},
	)
	return dat.(*wowgd.QuestCategory), header, err
}

// WoWQuestAreasIndex returns an index of quest areas.
func (c *Client) WoWQuestAreasIndex(ctx context.Context) (*wowgd.QuestAreasIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/quest/area/index",
		c.GetStaticNamespace(),
		&wowgd.QuestAreasIndex{},
	)
	return dat.(*wowgd.QuestAreasIndex), header, err
}

// WoWQuestArea returns a quest area by ID.
func (c *Client) WoWQuestArea(ctx context.Context, questAreaID int) (*wowgd.QuestArea, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/quest/area/%d", questAreaID),
		c.GetStaticNamespace(),
		&wowgd.QuestArea{},
	)
	return dat.(*wowgd.QuestArea), header, err
}

// WoWQuestTypesIndex returns an index of quest types (such as PvP quests, raid quests, or account quests).
func (c *Client) WoWQuestTypesIndex(ctx context.Context) (*wowgd.QuestTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/quest/type/index",
		c.GetStaticNamespace(),
		&wowgd.QuestTypesIndex{},
	)
	return dat.(*wowgd.QuestTypesIndex), header, err
}

// WoWQuestType returns a quest type by ID.
func (c *Client) WoWQuestType(ctx context.Context, questTypeID int) (*wowgd.QuestType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/quest/type/%d", questTypeID),
		c.GetStaticNamespace(),
		&wowgd.QuestType{},
	)
	return dat.(*wowgd.QuestType), header, err
}

// WoWRealmIndex returns an index of realms.
func (c *Client) WoWRealmIndex(ctx context.Context) (*wowgd.RealmIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/realm/index",
		c.GetDynamicNamespace(),
		&wowgd.RealmIndex{},
	)
	return dat.(*wowgd.RealmIndex), header, err
}

// WoWRealm returns a single realm by slug or ID.
func (c *Client) WoWRealm(ctx context.Context, realmSlug string) (*wowgd.Realm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/realm/%s", realmSlug),
		c.GetDynamicNamespace(),
		&wowgd.Realm{},
	)
	return dat.(*wowgd.Realm), header, err
}

// WoWRealmSearch searches for realms
func (c *Client) WoWRealmSearch(ctx context.Context, opts ...wowsearch.Opt) (*wowgd.RealmSearch, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/search/realm%s", buildSearchParams(opts...)),
		c.GetDynamicClassicNamespace(),
		&wowgd.RealmSearch{},
	)
	return dat.(*wowgd.RealmSearch), header, err
}

// WoWRegionIndex returns an index of regions.
func (c *Client) WoWRegionIndex(ctx context.Context) (*wowgd.RegionIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/region/index",
		c.GetDynamicNamespace(),
		&wowgd.RegionIndex{},
	)
	return dat.(*wowgd.RegionIndex), header, err
}

// WoWRegion returns a single region by ID.
func (c *Client) WoWRegion(ctx context.Context, regionID int) (*wowgd.Region, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/region/%d", regionID),
		c.GetDynamicNamespace(),
		&wowgd.Region{},
	)
	return dat.(*wowgd.Region), header, err
}

// WoWReputationFactionsIndex returns an index of reputation factions.
func (c *Client) WoWReputationFactionsIndex(ctx context.Context) (*wowgd.ReputationFactionsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/reputation-faction/index",
		c.GetStaticNamespace(),
		&wowgd.ReputationFactionsIndex{},
	)
	return dat.(*wowgd.ReputationFactionsIndex), header, err
}

// WoWReputationFaction returns a single reputation faction by ID.
func (c *Client) WoWReputationFaction(ctx context.Context, reputationFactionID int) (*wowgd.ReputationFaction, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-tier/%d", reputationFactionID),
		c.GetStaticNamespace(),
		&wowgd.ReputationFaction{},
	)
	return dat.(*wowgd.ReputationFaction), header, err
}

// WoWReputationTiersIndex returns an index of reputation tiers.
func (c *Client) WoWReputationTiersIndex(ctx context.Context) (*wowgd.ReputationTiersIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/reputation-tiers/index",
		c.GetStaticNamespace(),
		&wowgd.ReputationTiersIndex{},
	)
	return dat.(*wowgd.ReputationTiersIndex), header, err
}

// WoWReputationTiers returns a single set of reputation tiers by ID.
func (c *Client) WoWReputationTiers(ctx context.Context, reputationTiersID int) (*wowgd.ReputationTiers, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/reputation-tiers/%d", reputationTiersID),
		c.GetStaticNamespace(),
		&wowgd.ReputationTiers{},
	)
	return dat.(*wowgd.ReputationTiers), header, err
}

// WoWSpell returns a spell by ID.
func (c *Client) WoWSpell(ctx context.Context, spellID int) (*wowgd.Spell, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/spell/%d", spellID),
		c.GetStaticNamespace(),
		&wowgd.Spell{},
	)
	return dat.(*wowgd.Spell), header, err
}

// WoWSpellMedia returns media for a spell by ID.
func (c *Client) WoWSpellMedia(ctx context.Context, spellID int) (*wowgd.SpellMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/spell/%d", spellID),
		c.GetStaticNamespace(),
		&wowgd.SpellMedia{},
	)
	return dat.(*wowgd.SpellMedia), header, err
}

// WoWTalentsIndex returns an index of talents.
func (c *Client) WoWTalentsIndex(ctx context.Context) (*wowgd.TalentsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/talent/",
		c.GetStaticNamespace(),
		&wowgd.TalentsIndex{},
	)
	return dat.(*wowgd.TalentsIndex), header, err
}

// WoWTalent returns a talent by ID.
func (c *Client) WoWTalent(ctx context.Context, talentID int) (*wowgd.Talent, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/talent/%d", talentID),
		c.GetStaticNamespace(),
		&wowgd.Talent{},
	)
	return dat.(*wowgd.Talent), header, err
}

// WoWPvPTalentsIndex returns an index of PvP talents.
func (c *Client) WoWPvPTalentsIndex(ctx context.Context) (*wowgd.PvPTalentsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pvp-talent/",
		c.GetStaticNamespace(),
		&wowgd.PvPTalentsIndex{},
	)
	return dat.(*wowgd.PvPTalentsIndex), header, err
}

// WoWPvPTalent returns a PvP talent by ID.
func (c *Client) WoWPvPTalent(ctx context.Context, pvpTalentID int) (*wowgd.PvPTalent, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-talent/%d", pvpTalentID),
		c.GetStaticNamespace(),
		&wowgd.PvPTalent{},
	)
	return dat.(*wowgd.PvPTalent), header, err
}

// WoWTechTalentTreeIndex returns an index of tech talent trees.
func (c *Client) WoWTechTalentTreeIndex(ctx context.Context) (*wowgd.TechTalentTreeIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/tech-talent-tree/index",
		c.GetStaticNamespace(),
		&wowgd.TechTalentTreeIndex{},
	)
	return dat.(*wowgd.TechTalentTreeIndex), header, err
}

// WoWTechTalentTree returns a tech talent tree by ID.
func (c *Client) WoWTechTalentTree(ctx context.Context, techTalentTreeID int) (*wowgd.TechTalentTree, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/tech-talent-tree/%d", techTalentTreeID),
		c.GetStaticNamespace(),
		&wowgd.TechTalentTree{},
	)
	return dat.(*wowgd.TechTalentTree), header, err
}

// WoWTechTalentIndex returns an index of tech talents.
func (c *Client) WoWTechTalentIndex(ctx context.Context) (*wowgd.TechTalentIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/tech-talent/index",
		c.GetStaticNamespace(),
		&wowgd.TechTalentIndex{},
	)
	return dat.(*wowgd.TechTalentIndex), header, err
}

// WoWTechTalent returns a tech talent by ID.
func (c *Client) WoWTechTalent(ctx context.Context, techTalentID int) (*wowgd.TechTalent, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/tech-talent/%d", techTalentID),
		c.GetStaticNamespace(),
		&wowgd.TechTalent{},
	)
	return dat.(*wowgd.TechTalent), header, err
}

// WoWTechTalentMedia returns media for a tech talent by ID.
func (c *Client) WoWTechTalentMedia(ctx context.Context, techTalentID int) (*wowgd.TechTalentMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/tech-talent/%d", techTalentID),
		c.GetStaticNamespace(),
		&wowgd.TechTalentMedia{},
	)
	return dat.(*wowgd.TechTalentMedia), header, err
}

// WoWTitlesIndex returns an index of titles.
func (c *Client) WoWTitlesIndex(ctx context.Context) (*wowgd.TitlesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/title/index",
		c.GetStaticNamespace(),
		&wowgd.TitlesIndex{},
	)
	return dat.(*wowgd.TitlesIndex), header, err
}

// WoWTitle returns a title by ID.
func (c *Client) WoWTitle(ctx context.Context, titleID int) (*wowgd.Title, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/title/%d", titleID),
		c.GetStaticNamespace(),
		&wowgd.Title{},
	)
	return dat.(*wowgd.Title), header, err
}

// WoWToken returns the WoW Token index
func (c *Client) WoWToken(ctx context.Context) (*wowgd.Token, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/token/index",
		c.GetDynamicNamespace(),
		&wowgd.Token{},
	)
	return dat.(*wowgd.Token), header, err
}
