package definitions

func init() {
	add(`GroupDetails`, &defGroupDetails{})
}

type defGroupDetails struct{}

func (*defGroupDetails) String() string {
	return `<interface>
  <object class="GtkDialog" id="dialog">
    <property name="window-position">GTK_WIN_POS_CENTER</property>
    <property name="title" translatable="yes">Add Group</property>
    <child internal-child="vbox">
      <object class="GtkBox" id="Vbox">
        <property name="homogeneous">false</property>
        <property name="orientation">GTK_ORIENTATION_VERTICAL</property>
        <child>
          <object class="GtkLabel" id="group-name-label" >
            <property name="label" translatable="yes">Group name</property>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkEntry" id="group-name">
            <property name="has-focus">true</property>
            <property name="activates-default">True</property>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">2</property>
          </packing>
        </child>
        <child internal-child="action_area">
          <object class="GtkButtonBox" id="bbox">
            <property name="orientation">GTK_ORIENTATION_HORIZONTAL</property>
            <child>
              <object class="GtkButton" id="btn-cancel">
                <property name="label" translatable="yes">Cancel</property>
              </object>
            </child>
            <child>
              <object class="GtkButton" id="btn-ok">
                <property name="label" translatable="yes">OK</property>
                <property name="can-default">true</property>
              </object>
            </child>
          </object>
        </child>
      </object>
    </child>
    <action-widgets>
      <action-widget response="cancel">btn-cancel</action-widget>
      <action-widget response="ok" default="true">btn-ok</action-widget>
    </action-widgets>
  </object>
</interface>`
}
